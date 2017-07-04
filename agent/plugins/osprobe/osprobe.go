package osprobe

import (
	"fmt"
	"runtime"
	"os"
	"github.com/cloudfoundry/gosigar"
	"time"
	"git.oschina.net/k2ops/jarvis/agent/plugins/osprobe/disk"
)

type osInfo struct {
	UpdatedAt string
	Type string
	Arch string
	Hostname string
	CPUNum int
	MemTotal uint64
	UpTime string
}

func Detect() {
	osinfo := osInfo{}
	osinfo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05Z0700")
	osinfo.Type = runtime.GOOS
	osinfo.Arch = runtime.GOARCH
	osinfo.Hostname, _ = os.Hostname()
	osinfo.CPUNum = runtime.NumCPU()
	iSigar := sigar.ConcreteSigar{}
	memInfo, _ := iSigar.GetMem()
	osinfo.MemTotal = memInfo.Total
	uptime := sigar.Uptime{}
	uptime.Get()
	osinfo.UpTime = uptime.Format()

	fmt.Println(disk.PhysicalDisks())



	fmt.Println(osinfo)
}

