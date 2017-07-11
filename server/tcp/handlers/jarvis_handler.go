package handlers

import (
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"git.oschina.net/k2ops/jarvis/protocol"
	"net"
	"errors"
)

type JarvisHandler struct {

}
func (h *JarvisHandler) Handle(c []byte, conn net.Conn) error {
	log.Debug(string(c))
	m := &protocol.JarvisMessage{}
	err := json.Unmarshal(c, m)
	if err!=nil {
		log.Error(err.Error())
	}
	switch m.MessageType {
	case "heartbeat":
		return h.handleHeartbeat(c, conn)
	default:
		return errors.New("")
	}
}

func (*JarvisHandler) handleHeartbeat(c []byte, conn net.Conn) error {
	m := protocol.NewHeartbeatMessage()
	if err := json.Unmarshal(c, m); err!=nil {
		log.Error("can not unmarshal HeartbeatMessage ", err.Error())
		return err
	}
	log.Info(m)
	return nil
}
