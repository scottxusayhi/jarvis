package plugins

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"time"
)

func HeartBeat() {
	for {
		if core.Healthy() {
			msg := protocol.NewHeartbeatMessage(core.AgentId).Serialize()
			_, err := core.Conn.Write(msg)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Heartbeat send failed.")
			} else {
				core.LogMsgSent(msg)
			}
			time.Sleep(time.Duration(options.HBInterval) * time.Second)
		}
	}
}
