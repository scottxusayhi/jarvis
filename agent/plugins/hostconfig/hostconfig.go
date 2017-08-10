package hostconfig

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/agent/plugins/hostconfig/disk"
	"git.oschina.net/k2ops/jarvis/protocol"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)

func NewHostConfigMessage() *protocol.HostConfigMessage {
	m := protocol.NewEmptyHostConfigMessage(core.AgentId)
	// os info
	m.OsDetected.OsType = runtime.GOOS
	m.OsDetected.Arch = runtime.GOARCH
	m.OsDetected.Hostname, _ = os.Hostname()
	m.OsDetected.Uptime, _ = host.Uptime()

	// cpu info
	m.CpuDetected.Vcpu = runtime.NumCPU()
	// memory info
	memInfo, _ := mem.VirtualMemory()
	m.MemDetected.Total = memInfo.Total
	m.MemDetected.Available = memInfo.Available
	m.MemDetected.Used = memInfo.Used
	// disk info
	var err error
	m.DiskDetected, err = disk.PhysicalDisks()
	if err != nil {
		log.WithError(err).Error("detect disks failed")
	}
	return m
}

func Detect() {
	for {
		if core.Healthy() {
			m := NewHostConfigMessage()
			_, err := core.Conn.Write(m.Serialize())
			if err != nil {
				log.WithError(err).Error("Host config info send failed")
			} else {
				core.LogMsgSent(m.Serialize())
			}
			time.Sleep(30 * time.Second)
		}
	}
}
