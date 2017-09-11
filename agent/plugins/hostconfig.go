package plugins

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/protocol"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"time"
	"git.oschina.net/k2ops/jarvis/agent/plugins/hostconfig"
)

func NewHostConfigMessage() *protocol.HostConfigMessage {
	m := protocol.NewEmptyHostConfigMessage(core.AgentId)
	// os info
	m.OsDetected = hostconfig.OsInfo()

	// cpu info
	m.CpuDetected = hostconfig.CpuInfo()
	// memory info
	memInfo, _ := mem.VirtualMemory()
	m.MemDetected.Total = memInfo.Total
	m.MemDetected.Available = memInfo.Available
	m.MemDetected.Used = memInfo.Used
	// disk info
	var err error
	m.DiskDetected, err = hostconfig.PhysicalDisks()
	if err != nil {
		log.WithError(err).Error("detect disks failed")
	}
	// network
	m.NetworkDetected.Ip = hostconfig.ExternalIP()
	return m
}

func HostConfig() {
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
