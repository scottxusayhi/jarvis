package tcp

import (
	"net"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
	"bufio"
	"io"
	"git.oschina.net/k2ops/jarvis/utils"
	"errors"
)

type JarvisHandler struct {
	conn net.Conn
	agentId string
	reader *bufio.Reader
}

func NewJarvisHandler(conn net.Conn) *JarvisHandler {
	return &JarvisHandler{
		conn: conn,
		reader: bufio.NewReader(conn),
	}
}

func (h *JarvisHandler ) Start () {
	go func() {
		log.WithFields(log.Fields{
			"localAddr": h.conn.LocalAddr(),
			"remoteAddr": h.conn.RemoteAddr(),
		}).Info("New Agent Connection")
		defer h.conn.Close()

		// send welcome message
		h.sendWelcome()

		for {
			// read
			raw, err := h.reader.ReadBytes(protocol.Footer)
			if err == io.EOF {
				log.WithFields(log.Fields{
					"remoteAddr": h.conn.RemoteAddr(),
				}).Info("Connection Closed")
				break
			} else if err != nil {
				log.Error(err.Error())
			} else {
				h.LogMsgReceived(raw)
				h.handleMessage(raw)
			}
		}
	}()
}

func (h *JarvisHandler) handleMessage(raw []byte) error {
	msgType, err := protocol.MsgType(raw)
	if err != nil {
		return err
	}
	switch msgType {
	case protocol.MSG_AGENT_ID_REQUEST:
		return h.handleAgentIdRequest(raw)
		break
	default:
		return errors.New("unknown message type " + msgType)
	}
	return nil
}

func (h *JarvisHandler) handleAgentIdRequest(raw []byte) error {
	return h.sendAgentIdResponse()
}

func (h *JarvisHandler ) sendWelcome() error {
	msg := protocol.NewWelcomeMessage(h.conn.RemoteAddr().String(), h.conn.LocalAddr().String()).Serialize()
	h.LogMsgSent(msg)
	_, err := h.conn.Write(msg)
	return err
}

func (h *JarvisHandler) sendAgentIdResponse() error {
	msg := protocol.NewAgentIdResponse("this is your id").Serialize()
	h.LogMsgSent(msg)
	_, err := h.conn.Write(msg)
	return err
}

func (h *JarvisHandler) LogMsgSent(msg []byte) {
	utils.LogMsgSent(msg, h.agentId)
}

func (h *JarvisHandler) LogMsgReceived(msg []byte) {
	utils.LogMsgReceived(msg, h.agentId)
}
