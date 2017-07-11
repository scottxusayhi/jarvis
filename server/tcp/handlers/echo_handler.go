package handlers

import (
	log "github.com/sirupsen/logrus"
	"net"
)



type EchoHandler struct {}
func (h *EchoHandler) Handle(c []byte, conn net.Conn) error {
	log.Debug(string(c))
	_, err := conn.Write(c)
	return err
}