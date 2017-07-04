package osprobe

import (
	"fmt"
	"runtime"
	"os"
	"time"
	"git.oschina.net/k2ops/jarvis/agent/plugins/osprobe/disk"
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
	"git.oschina.net/k2ops/jarvis/utils"
)

type osInfo struct {
	UpdatedAt string
	Type string
	Arch string
	Hostname string
	CPUNum int
	MemTotal uint64
	UpTime string
	Disks []disk.PhysicalDiskInfo
}

func Detect() {
	osinfo := osInfo{}
	// updated at
	osinfo.UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z0700")
	// os type
	osinfo.Type = runtime.GOOS
	// os arch
	osinfo.Arch = runtime.GOARCH
	// hostname
	osinfo.Hostname, _ = os.Hostname()
	// cpu number
	osinfo.CPUNum = runtime.NumCPU()
	// total memory
	memInfo, _ := mem.VirtualMemory()
	osinfo.MemTotal = memInfo.Total
	// uptime
	uptime, _ := host.Uptime()
	osinfo.UpTime = utils.FormatSecond(uptime)
	// physical disks
	osinfo.Disks, _ = disk.PhysicalDisks()

	// print json
	jsonBytes, _ := json.Marshal(osinfo)
	fmt.Println(string(jsonBytes))
}

