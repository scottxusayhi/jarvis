package core

import (
	"bufio"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

var (
	Conn    net.Conn
	Reader  *bufio.Reader
	AgentId string

	Connected bool
	HasId     bool
)

func KeepConnected() {
	for ; ; time.Sleep(10 * time.Second) {
		if !Connected {
			connect()
			sayHello()
		}
		if !HasId {
			negotiateAgentId()
		}
	}
}

func connect() {
	var err error
	for ; ; time.Sleep(10 * time.Second) {
		log.WithFields(log.Fields{
			"master": options.Master,
		}).Debug("Trying connect to master")
		Conn, err = net.DialTimeout("tcp", options.Master, 3*time.Second)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		log.WithFields(log.Fields{
			"localAddr":  Conn.LocalAddr().String(),
			"remoteAddr": Conn.RemoteAddr().String(),
		}).Info("Connected")
		Reader = bufio.NewReader(Conn)
		break
	}
	Connected = true
}

func negotiateAgentId() {
	agentId, err := options.ReadAgentIdFromFile()
	if err != nil {
		log.Info("No agent id file found, request master for id")
		sendAgentIdRequest()
		// the response will be handled in another routine
	} else {
		AgentId = agentId
		HasId = true
	}
}

func sayHello() error {
	msg := protocol.NewHelloMessage().Serialize()
	_, err := Conn.Write(msg)
	LogMsgSent(msg)
	return err
}

func sendAgentIdRequest() error {
	msg := protocol.NewAgentIdRequest().Serialize()
	_, err := Conn.Write(msg)
	LogMsgSent(msg)
	return err
}

func Healthy() bool {
	return Connected && HasId
}

func LogMsgSent(msg []byte) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v -> sent message", AgentId)
}

func LogMsgReceived(msg []byte) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v <- received message", AgentId)
}

func UpdateAgentId(id string) error {
	if err := options.WriteBackAgentIdFile(id); err != nil {
		return err
	}
	AgentId = id
	HasId = true
	return nil
}
