package main

import (
	"github.com/scottxusayhi/jarvis/agent/core"
	"github.com/scottxusayhi/jarvis/agent/options"
	"github.com/scottxusayhi/jarvis/agent/plugins"
	"github.com/scottxusayhi/jarvis/utils"
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
