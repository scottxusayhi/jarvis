package plugins

import (
	"github.com/scottxusayhi/jarvis/agent/core"
	"github.com/scottxusayhi/jarvis/protocol"
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
