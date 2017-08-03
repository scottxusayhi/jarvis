package conn

import (
	"net"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"time"
	"bufio"
)

var (
	Conn    net.Conn
	Reader  *bufio.Reader
	AgentId string

	Connected bool
	HasId     bool
)


func KeepConnected() {
	for ;;time.Sleep(10*time.Second) {
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
	for ;;time.Sleep(10*time.Second){
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
	agentId, err := options.GetAgentIdFromFile()
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

