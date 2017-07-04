// +build darwin

package disk

import (
	"github.com/shirou/gopsutil/disk"
	"fmt"
	"regexp"
)

type PhysicalDiskInfo struct {
	Device string `json:"device"`
	Total  uint64 `json:"total"`
	Used uint64 `json:"used"`
}

type PartitionInfoWithCapa struct {
	MountInfo disk.PartitionStat
	Total uint64
	Used uint64
}

func getDiskName(device string) string {
	myExp, _ := regexp.Compile("/dev/disk(?P<diskIndex>\\d)s(?P<partIndex>\\d)")
	match := myExp.FindStringSubmatch(device)
	result := make(map[string]string)
	for i, name := range myExp.SubexpNames() {
		if i!=0 {
			result[name] = match[i]
		}
	}
	return fmt.Sprintf("/dev/disk%d", result["diskIndex"])
}

func findPDInfo(info []PhysicalDiskInfo, diskIndex string) int {
	for index, pdInfo := range info {
		if pdInfo.Device == diskIndex {
			return index
		}
	}
	return -1
}

func aggregatePartInfo(info []PhysicalDiskInfo, device string, total uint64, used uint64) []PhysicalDiskInfo {
	//fmt.Println("====")
	fmt.Println(info, device, total, used)
	diskIndex := getDiskName(device)
	pdInfo := findPDInfo(info, diskIndex)
	if pdInfo == -1 {
		// new disk
		//fmt.Println("new disk")
		d := PhysicalDiskInfo{
			Device: diskIndex,
			Total:  total,
			Used: used,
		}
		info = append(info, d)
	} else {
		//fmt.Println("merge partition to disk")
		//fmt.Printf("before: %v\n", info[pdInfo])
		d := PhysicalDiskInfo{
			Device: diskIndex,
			Total:  total+info[pdInfo].Total,
			Used: used+info[pdInfo].Used,
		}
		info = append(info, d)
		info = append(info[:pdInfo], info[pdInfo+1])
		//fmt.Printf("after: %v\n", info[pdInfo])
	}
	return info
}


func PhysicalDisks() ([]PhysicalDiskInfo, error) {
	// return value
	var ret []PhysicalDiskInfo
	// gather disk partition with usage info
	var myPartInfo []PartitionInfoWithCapa
	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		//fmt.Println(index, part)
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return ret, err
		}
		//fmt.Println(usage)

		p := PartitionInfoWithCapa{
			MountInfo: part,
			Total: usage.Total,
			Used: usage.Used,
		}
		myPartInfo = append(myPartInfo, p)
	}

	// group partitions by disk
	for _, myPart := range myPartInfo {
		ret = aggregatePartInfo(ret, myPart.MountInfo.Device, myPart.Total, myPart.Used)
	}
	//fmt.Println(ret)
	return ret, nil
}
