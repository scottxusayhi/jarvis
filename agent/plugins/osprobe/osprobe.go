package osprobe

import (
	"encoding/json"
	"fmt"
	"git.oschina.net/k2ops/jarvis/agent/plugins/osprobe/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"runtime"
	"time"
	"git.oschina.net/k2ops/jarvis/utils"
)

type osInfo struct {
	UpdatedAt string
	Type      string
	Arch      string
	Hostname  string
	CPUNum    int
	MemTotal  uint64
	UpTime    string
	Disks     []disk.PhysicalDiskInfo
}

func Detect() {
	osInfo := osInfo{}
	// updated at
	osInfo.UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z0700")
	// os type
	osInfo.Type = runtime.GOOS
	// os arch
	osInfo.Arch = runtime.GOARCH
	// hostname
	osInfo.Hostname, _ = os.Hostname()
	// cpu number
	osInfo.CPUNum = runtime.NumCPU()
	// total memory
	memInfo, _ := mem.VirtualMemory()
	osInfo.MemTotal = memInfo.Total
	// up time
	upTime, _ := host.Uptime()
	osInfo.UpTime = timeutils.FormatSecond(upTime)
	// physical disks
	osInfo.Disks, _ = disk.PhysicalDisks()

	// print json
	jsonBytes, _ := json.Marshal(osInfo)
	fmt.Println(string(jsonBytes))
}
