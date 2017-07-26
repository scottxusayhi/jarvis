package tcp

import (
	"net"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
)

type JarvisMaster struct {
	conn net.Conn
}

func (c *JarvisMaster) SendWelcome() {
	_, err := c.conn.Write(protocol.NewWelcomeMessage(c.conn.RemoteAddr().String(), c.conn.LocalAddr().String()).Serialize())
	if err!=nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to send welcome message")
	}
}

func NewJarvisMaster(conn net.Conn) *JarvisMaster {
	return &JarvisMaster{
		conn: conn,
	}
}
