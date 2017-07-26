package main

import (
	"net"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
)

type JarvisAgent struct {
	conn net.Conn
}

func NewJarvisAgent (conn net.Conn) *JarvisAgent {
	return &JarvisAgent{
		conn: conn,
	}
}

func (c *JarvisAgent) SendHeartbeatMessage() {
	_, err := c.conn.Write(protocol.NewHeartbeatMessage().Serialize())
	if err!=nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to send heartbeat message")
	}
}

func (a *JarvisAgent) HandleWelcome ()  {
	log.Info("response to welcome")
}

func (a *JarvisAgent) HandleMetadataChange () {
	// TODO: update agent metadata (datacenter, rack and slot) and save back to config file so it can take effect after next restart
	log.Info("response to metadata change")
}

