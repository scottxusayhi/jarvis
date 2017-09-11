// +build darwin

package hostconfig

import (
	"fmt"
	"git.oschina.net/k2ops/jarvis/protocol"
	"github.com/shirou/gopsutil/disk"
	"regexp"
	log "github.com/sirupsen/logrus"
)

// extract disk device name (e.g. /dev/disk1) from partition name
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

func index(info protocol.HostDisks, diskName string) int {
	for index, pdInfo := range info {
		if pdInfo.Device == diskName {
			return index
		}
	}
	return -1
}

// aggregate partition info into disk (s)
func aggregatePartInfo(info protocol.HostDisks, device string, cap uint64, used uint64) protocol.HostDisks {
	//fmt.Println("====")
	//fmt.Println(info, device, total, used)
	pdName := getDiskName(device)
	idx := index(info, pdName)
	if idx == -1 {
		// new disk
		d := protocol.DiskInfo{
			Device:   pdName,
			Capacity: cap,
			Used:     used,
		}
		info = append(info, d)
	} else {
		//fmt.Println("merge partition to disk")
		//fmt.Printf("before: %v\n", info[pdInfo])
		d := protocol.DiskInfo{
			Device:   pdName,
			Capacity: cap + info[idx].Capacity,
			Used:     used + info[idx].Used,
		}
		// append to tail
		info = append(info, d)
		// slice trick: delete an element
		info = append(info[:idx], info[idx+1])
		//fmt.Printf("after: %v\n", info[pdInfo])
	}
	return info
}

func PhysicalDisks() (disks protocol.HostDisks, err error) {
	parts, _ := disk.Partitions(false)
	for _, part := range parts {
		log.Debug(part)
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return disks, err
		}
		//fmt.Println(usage)
		// group partitions to disk
		disks = aggregatePartInfo(disks, part.Device, usage.Total, usage.Used)
	}
	return disks, nil
}
