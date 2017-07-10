package main

import (
	"net"
	"fmt"
	"time"
	"io"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/utils"
	"git.oschina.net/k2ops/jarvis/agent/options"
	log "github.com/sirupsen/logrus"
)

func heartBeat(conn net.Conn, interval time.Duration) {
	for {
		_, err := conn.Write(protocol.NewHeartbeatMessage().Serialize())
		if err!=nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Heartbeat send failed.")
		} else {
			log.Info("Heartbeat sent.")
		}
		time.Sleep(interval*time.Second)
	}
}

func Listen(conn net.Conn) error {
	response := make([]byte, 100)
	for {
		// reset buffer
		for index := range response {
			response[index] = 0
		}
		n, err := conn.Read(response)
		if err == io.EOF {
			log.Error("Connection closed by remote")
			return err
		} else if err != nil {
			fmt.Println(err.Error())
			return err
		}
		log.Info(string(response[:n]))
	}
}

func channel_hold() {
	c := make(chan string)
	for {
		log := <- c
		fmt.Println(log)
	}
}

func main() {
	// logger
	utils.InitLogger(log.InfoLevel)
	// CLI options
	options.Check()
	if options.Flags().Debug {
		log.SetLevel(log.DebugLevel)
		log.Info("Debug enabled.")
	}
	// connect
	conn, err := net.DialTimeout("tcp", options.Flags().Master, 3*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	log.WithFields(log.Fields{
		"localAddr": conn.LocalAddr().String(),
		"remoteAddr": conn.RemoteAddr().String(),
	}).Info("Connected.")

	// heart beat
	go heartBeat(conn, time.Duration(options.Flags().HBInterval))
	go Listen(conn)

	//KeyboardInput(conn)
	channel_hold()

}
