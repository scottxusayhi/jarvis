package main

import (
	"net"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"time"
	"bufio"
	"io"
	"fmt"
	"encoding/json"
	"errors"
)

var (
	connection net.Conn
	reader *bufio.Reader
	AgentId    string
)

func Connect() {
	var err error
	for ;;time.Sleep(10*time.Second){
		connection, err = net.DialTimeout("tcp", options.Master, 3*time.Second)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		log.WithFields(log.Fields{
			"localAddr":  connection.LocalAddr().String(),
			"remoteAddr": connection.RemoteAddr().String(),
		}).Info("Connected")
		reader = bufio.NewReader(connection)
		break
	}
	// agent should send something to trigger cmux to work, so cmux can route this connection to tcp server
	sayHello()
}

func readNextMessage() ([]byte, error) {
	raw, err := reader.ReadBytes(protocol.Footer)
	if err == io.EOF {
		log.Error("Connection closed by remote, try re-connect")
		Connect()
	} else if err != nil {
		log.Error(err.Error())
	}
	return raw, err
}

func NegotiateAgentId() {

}

func HeartBeat() {
	for {
		_, err := connection.Write(protocol.NewHeartbeatMessage().Serialize())
		if err!=nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Heartbeat send failed.")
		} else {
			log.Info("Heartbeat sent.")
		}
		time.Sleep(time.Duration(options.HBInterval)*time.Second)
	}
}

func DoMyJob() {
	reader := bufio.NewReader(connection)
	for {
		raw, err := reader.ReadBytes(protocol.Footer)
		if err == io.EOF {
			log.Error("Connection closed by remote")
		} else if err != nil {
			fmt.Println(err.Error())
		}
		handleMessage(raw)
	}
}

type jsonObject map[string]interface{}

func msgType(raw []byte) (string, error) {
	var err error
	msg := jsonObject{}
	err = json.Unmarshal(raw, &msg)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	msgType, ok := msg["type"].(string)
	if ok {
		return msgType, nil
	} else {
		return "", errors.New(fmt.Sprintf("msg type: expect string but got %T", msg["type"]))
	}
}

func handleMessage(raw []byte) error {
	msgType, err := msgType(raw)
	if err != nil {
		return err
	}
	switch msgType {
	case protocol.MSG_WELCOME:
		return handleWelcome(raw)
	case protocol.MSG_AGENT_ID_RESPONSE:
		return handleAgentIdResponse(raw)
	default:
		return errors.New("unknown message type " + msgType)
	}
}

func handleWelcome(raw []byte) error {
	log.WithFields(log.Fields{
		"msg": string(raw),
	}).Info("Received welcome")
	return nil
}

func handleAgentIdResponse(raw []byte) error {
	return nil
}

func sayHello() {
	_, err := connection.Write(protocol.NewHelloMessage().Serialize())
	if err != nil {
		log.Error(err.Error())
	}
}

