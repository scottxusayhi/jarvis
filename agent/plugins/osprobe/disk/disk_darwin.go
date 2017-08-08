// +build darwin

package disk

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"regexp"
)

func getDiskName(device string) string {
	myExp, _ := regexp.Compile("/dev/disk(?P<diskIndex>\\d)s(?P<partIndex>\\d)")
	match := myExp.FindStringSubmatch(device)
	result := make(map[string]string)
	for i, name := range myExp.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	return fmt.Sprintf("/dev/disk%v", result["diskIndex"])
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
	//fmt.Println(info, device, total, used)
	pdName := getDiskName(device)
	pdInfo := findPDInfo(info, pdName)
	if pdInfo == -1 {
		// new disk
		//fmt.Println("new disk")
		d := PhysicalDiskInfo{
			Device: pdName,
			Total:  total,
			Used:   used,
		}
		info = append(info, d)
	} else {
		//fmt.Println("merge partition to disk")
		//fmt.Printf("before: %v\n", info[pdInfo])
		d := PhysicalDiskInfo{
			Device: pdName,
			Total:  total + info[pdInfo].Total,
			Used:   used + info[pdInfo].Used,
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
	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		//fmt.Println(index, part)
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return ret, err
		}
		//fmt.Println(usage)
		// group partitions to disk
		ret = aggregatePartInfo(ret, part.Device, usage.Total, usage.Used)
	}
	return ret, nil
}
