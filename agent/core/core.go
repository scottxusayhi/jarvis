package core

import (
	"bufio"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/protocol"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
	"fmt"
)

var (
	Conn    net.Conn
	Reader  *bufio.Reader
	AgentId string

	Connected bool
	HasId     bool
)

func KeepConnected(con chan bool, id chan bool) {
	for ; ; time.Sleep(10 * time.Second) {
		log.Info(fmt.Sprintf("connected is %s", Connected))
		if !Connected {
			connect()
		} else {
			log.Info(fmt.Sprintf("connected is %s", Connected))
			con <- true
			log.Info(fmt.Sprintf("connected is %s", Connected))
			sayHello()
			log.Info(fmt.Sprintf("hasid is %s", HasId))
			if !HasId {
			        negotiateAgentId()
			} else {
				id  <- true
			}
		}
	}
}

func connect() {
	retryInterval := 10 * time.Second
	log.WithFields(log.Fields{
		"master": options.Master,
	}).Info("Trying connect to master")
	var err error
	Conn, err = net.DialTimeout("tcp", options.Master, 3*time.Second)
	if err != nil {
		log.WithError(err).Error(fmt.Sprintf("tcp connect failed, retry in %v", retryInterval))
	} else {
	    log.WithFields(log.Fields{
		    "localAddr":  Conn.LocalAddr().String(),
		    "remoteAddr": Conn.RemoteAddr().String(),
	    }).Info("Connected")
	    Reader = bufio.NewReader(Conn)
	    Connected = true
    }
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
	log.Info(fmt.Sprintf("hasid is %s", HasId))
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
