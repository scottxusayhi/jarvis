package plugins

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
)

func HeartBeat() {
	msg := protocol.NewHeartbeatMessage(core.AgentId).Serialize()
	_, err := core.Conn.Write(msg)
	if err != nil {
		log.WithFields(log.Fields{
		"error": err.Error(),
		}).Error("Heartbeat send failed.")
	} else {
		core.LogMsgSent(msg)
	}
}
