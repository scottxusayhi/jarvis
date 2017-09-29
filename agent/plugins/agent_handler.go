package plugins

import (
	"encoding/json"
	"errors"
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

func HandleMsg() {
	for {
		if core.Reader != nil {
			// if no data in reader, program will be blocked
			raw, err := core.Reader.ReadBytes(protocol.Footer)
			if err == io.EOF {
				log.Error("Connection closed by remote")
				core.Connected = false
			} else if err != nil {
				log.Error(err.Error())
			}
			core.LogMsgReceived(raw)
			err = handleMessage(raw)
			if err != nil {
				log.Error(err.Error())
			}
		} else {
			time.Sleep(10*time.Second)
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
		//break
	case protocol.MSG_AGENT_ID_RESPONSE:
		return handleAgentIdResponse(raw)
		//break
	default:
		return errors.New("unknown message type " + msgType)
	}
	//return nil
}

func handleWelcome(raw []byte) error {
	return nil
}

func handleAgentIdResponse(raw []byte) error {
	var err error
	r := protocol.AgentIdResponse{}
	if err = json.Unmarshal(raw, &r); err != nil {
		return err
	}
	if err = core.UpdateAgentId(r.AgentId); err != nil {
		return err
	}
	return nil
}
