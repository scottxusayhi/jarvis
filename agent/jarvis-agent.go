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
		log.Debug("Debug enabled.")
	}
}

func report(id chan bool){
    for {
        select{
        case <- id:
            plugins.HeartBeat()
            plugins.HostConfig()
         }
   }
}

func main() {
	id := make(chan bool, 1)
	// options
	options.LoadCli()
	// logger
	initLogger()
	// connect
	go core.KeepConnected(id)
	go report(id)
	plugins.HandleMsg()
}
