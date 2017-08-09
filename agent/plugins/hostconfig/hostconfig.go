package hostconfig

import (
	"git.oschina.net/k2ops/jarvis/agent/plugins/hostconfig/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"runtime"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/agent/core"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewHostConfigMessage() *protocol.HostConfigMessage {
	m := protocol.NewEmptyHostConfigMessage()
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
	m.DiskDetected, _ = disk.PhysicalDisks()
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
			time.Sleep(30*time.Second)
		}
	}
}

