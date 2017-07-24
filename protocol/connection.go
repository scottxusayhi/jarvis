package protocol

import (
	log "github.com/sirupsen/logrus"
	"net"
)

type jarvisConn struct {
	conn net.Conn
}

func (c *jarvisConn) SendWelcome() {
	msg := append(NewWelcomeMessage(c.conn.RemoteAddr().String(), c.conn.LocalAddr().String()).Serialize(), Footer)
	_, err := c.conn.Write(msg)
	if err!=nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to send welcome message")
	}
}

func (c *jarvisConn) SendHeartbeatMessage() {
	msg := append(NewHeartbeatMessage().Serialize(), Footer)
	_, err := c.conn.Write(msg)
	if err!=nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to send heartbeat message")
	}
}

func NewConn(c net.Conn) jarvisConn {
	result := jarvisConn{}
	result.conn = c
	return result
}