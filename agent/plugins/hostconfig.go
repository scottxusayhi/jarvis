package plugins

import (
	"github.com/scottxusayhi/jarvis/agent/core"
	"github.com/scottxusayhi/jarvis/protocol"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"github.com/scottxusayhi/jarvis/agent/plugins/hostconfig"
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
	m := NewHostConfigMessage()
	_, err := core.Conn.Write(m.Serialize())
	if err != nil {
		log.WithError(err).Error("Host config info send failed")
	} else {
		core.LogMsgSent(m.Serialize())
	}
}
