package plugins

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	"io"
	log "github.com/sirupsen/logrus"
	"errors"
	"git.oschina.net/k2ops/jarvis/agent/conn"
)

func HandleMsg() {
	for {
		if conn.Connected {
			raw, err := conn.Reader.ReadBytes(protocol.Footer)
			if err == io.EOF {
				log.Error("Connection closed by remote")
				conn.Connected = false
			} else if err != nil {
				log.Error(err.Error())
			}
			conn.LogMsgReceived(raw)
			handleMessage(raw)
		}
	}
}

func handleMessage(raw []byte) error {
	msgType, err := protocol.MsgType(raw)
	if err != nil {
		return err
	}
	switch msgType {
	case protocol.MSG_WELCOME:
		return handleWelcome(raw)
		break
	case protocol.MSG_AGENT_ID_RESPONSE:
		return handleAgentIdResponse(raw)
		break
	default:
		return errors.New("unknown message type " + msgType)
	}
	return nil
}

func handleWelcome(raw []byte) error {
	return nil
}

func handleAgentIdResponse(raw []byte) error {
	return nil
}
