package plugins

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"time"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/agent/conn"
)

func HeartBeat() {
	for {
		if conn.Healthy() {
			_, err := conn.Conn.Write(protocol.NewHeartbeatMessage(conn.AgentId).Serialize())
			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Heartbeat send failed.")
			} else {
				log.Info("Heartbeat sent.")
			}
			time.Sleep(time.Duration(options.HBInterval) * time.Second)
		}
	}
}
