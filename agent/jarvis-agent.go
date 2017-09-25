package main

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/agent/plugins"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
    "time"
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
            time.Sleep(time.Duration(options.HBInterval - 10) * time.Second)
            plugins.HeartBeat()
            plugins.HostConfig()
         }
   }
}

func main() {
	connect := make(chan bool, 1)
	id := make(chan bool, 1)
	// options
	options.LoadCli()
	// logger
	initLogger()
	// connect
	go plugins.HandleMsg(connect)
	go report(id)
	core.KeepConnected(connect, id)
}
