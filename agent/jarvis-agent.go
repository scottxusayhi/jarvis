package main

import (
	"git.oschina.net/k2ops/jarvis/utils"
	"git.oschina.net/k2ops/jarvis/agent/options"
	log "github.com/sirupsen/logrus"
)


func initLogger() {
	utils.InitLogger(log.InfoLevel)
	if options.Debug {
		log.SetLevel(log.DebugLevel)
		log.Info("Debug enabled.")
	}
}

func main() {
	// options
	options.LoadCli()
	// logger
	initLogger()
	// connect
	Connect()
	NegotiateAgentId()
	go DoMyJob()
	HeartBeat()
}
