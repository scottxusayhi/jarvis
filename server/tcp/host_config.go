package tcp

import (
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
)

func Match(msg *protocol.HostConfigMessage) bool {
	backend, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		return false
	}
	osExpected, cpuExpected, memExpected, diskExpected, networkExpected, err := backend.QueryExpectedConfig(msg.AgentId)
	return osMatch(osExpected, &msg.OsDetected) && cpuMatch(cpuExpected, &msg.CpuDetected) && memMatch(memExpected, &msg.MemDetected) && diskMatch(diskExpected, &msg.DiskDetected) && networkMatch(networkExpected, &msg.NetworkDetected)
}

func osMatch(expected *model.OsInfo, detected *protocol.OsInfo) bool {
	if expected.Arch!=detected.Arch {
		return false
	}
	if expected.Hostname!=detected.Hostname {
		return false
	}
	if expected.OsType!=detected.OsType {
		return false
	}
	return true
}

func cpuMatch(expected *model.CpuInfo, detected *protocol.CpuInfo) bool {
	return true
}

func memMatch(expected *model.MemInfo, detected *protocol.MemInfo) bool {
	return true
}

func diskMatch(expected *model.HostDisks, detected *protocol.HostDisks) bool {
	return true
}

func networkMatch(expected *model.NetworkInfo, detected *protocol.NetworkInfo) bool {
	return true
}
