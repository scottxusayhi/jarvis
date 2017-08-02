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
)

var (
	Connection net.Conn
	AgentId string
)

func Connect() {
	var err error
	Connection, err = net.DialTimeout("tcp", options.Master, 3*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithFields(log.Fields{
		"localAddr": Connection.LocalAddr().String(),
		"remoteAddr": Connection.RemoteAddr().String(),
	}).Info("Connected.")
}

func WaitForWelcome() {

}

func NegotiateAgentId() {

}

func HeartBeat() {
	for {
		_, err := Connection.Write(protocol.NewHeartbeatMessage().Serialize())
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
	reader := bufio.NewReader(Connection)
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
func handleMessage(raw []byte) {
	var err error
	msg := jsonObject{}
	err = json.Unmarshal(raw, msg)
	if err != nil {
		log.Error(err.Error())
	}
	msgType, ok := msg["type"].(string)
	if !ok {
		log.Error("can not parse message type: not a string")
	}
	switch msgType {
	case protocol.MSG_WELCOME:
		handleWelcome(raw)
	}

}

func handleWelcome(raw []byte) {

}

func handleAgentIdResponse(raw []byte) {

}

