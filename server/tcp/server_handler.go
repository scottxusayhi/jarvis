package tcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/scottxusayhi/jarvis/protocol"
	"github.com/scottxusayhi/jarvis/server/backend/mysql"
	"github.com/scottxusayhi/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strconv"
	"time"
	"github.com/scottxusayhi/jarvis/server/options"
)

type JarvisHandler struct {
	conn          net.Conn
	agentId       int64 // init by heartbeat
	reader        *bufio.Reader
}

func NewJarvisHandler(conn net.Conn) *JarvisHandler {
	return &JarvisHandler{
		conn:          conn,
		reader:        bufio.NewReader(conn),
	}
}

func (h *JarvisHandler) Start() {
	go func() {
		log.WithFields(log.Fields{
			"localAddr":  h.conn.LocalAddr(),
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
				backend, err := mysql.GetBackend()
				if err != nil {
					log.Error(err.Error())
				}
				backend.MarkOffline(h.agentId)
				break
			} else if err != nil {
				log.Error(err.Error())
			} else {
				h.LogMsgReceived(raw)
				err = h.handleMessage(raw)
				if err != nil {
					log.Error(err.Error())
				}
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
	case protocol.MSG_HELLO:
		break
	case protocol.MSG_HEARTBEAT:
		return h.handleHeartBeat(raw)
		break
	case protocol.MSG_HOST_CONFIG:
		return h.handleHostConfig(raw)
		break
	default:
		return errors.New("unknown message type " + msgType)
	}
	return nil
}

// handlers
// agent id request
func (h *JarvisHandler) handleAgentIdRequest(raw []byte) error {
	backend, err := mysql.GetBackend()
	if err != nil {
		return err
	}
	newId, err := backend.PreserveId()
	if err != nil {
		return err
	}
	err = backend.UpdateConnectionInfo(newId)
	if err != nil {
		return err
	}
	return h.sendAgentIdResponse(newId)
}

// heartbeat
func (h *JarvisHandler) handleHeartBeat(raw []byte) error {
	hb := protocol.HeartbeatMessage{}
	err := json.Unmarshal(raw, &hb)
	if err != nil {
		return err
	}

	// init h.agentId for the first time
	id, err := strconv.ParseInt(hb.AgentId, 10, 64)
	if err != nil {
		return err
	}
	if h.agentId == 0 {
		h.agentId = id
	}

	// check whether id matches
	if h.agentId != id {
		return errors.New(fmt.Sprintf("agent id does not match: expect %v but got %v", h.agentId, id))
	}

	// update heartbeat
	backend, err := mysql.GetBackend()
	if err != nil {
		return err
	}
	hbTime := hb.UpdatedAt
	if options.UseMasterTime {
		hbTime = time.Now()
	}
	log.Debug(fmt.Sprintf("use-master-time is %s",options.UseMasterTime))
	backend.UpdateHeartBeat(h.agentId, hbTime)
	return nil
}

// host config
func (h *JarvisHandler) handleHostConfig(raw []byte) (err error)  {
	msg := protocol.HostConfigMessage{}
	if err = json.Unmarshal(raw, &msg); err!=nil {
		return err
	}

	// get mysql backend
	backend, err := mysql.GetBackend()
	if err!=nil {
		return err
	}

	// update database
	err = backend.UpdateHostConfig(
		intId(msg.AgentId),
		msg.OsDetected,
		msg.CpuDetected,
		msg.MemDetected,
		msg.DiskDetected,
		msg.NetworkDetected,
		Match(&msg))
	return err
}

// message to send
func (h *JarvisHandler) sendWelcome() error {
	msg := protocol.NewWelcomeMessage(h.conn.RemoteAddr().String(), h.conn.LocalAddr().String()).Serialize()
	h.LogMsgSent(msg)
	_, err := h.conn.Write(msg)
	return err
}

func (h *JarvisHandler) sendAgentIdResponse(newId int64) error {
	msg := protocol.NewAgentIdResponse(fmt.Sprintf("%v", newId)).Serialize()
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


// id converter
func intId(id string) (intId int64) {
	intId, _ = strconv.ParseInt(id, 10, 64)
	return intId
}

func strId(id int64) (string) {
	return fmt.Sprintf("%v", id)
}

// online check
func OnlineCheck() {
	backend, err := mysql.GetBackend()
	if err != nil {
		log.WithError(err).Error("OnlineCheck get backend")
	}
	backend.GrimReaper()
}
