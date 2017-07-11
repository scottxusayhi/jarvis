package handlers

import (
	"net"
)

type MsgHandler interface {
	Handle(c []byte, conn net.Conn) error
}


