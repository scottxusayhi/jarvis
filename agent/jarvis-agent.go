package main

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/agent/plugins"
	"git.oschina.net/k2ops/jarvis/utils"
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
	go core.KeepConnected()
	go plugins.HeartBeat()
	plugins.HandleMsg()
}
