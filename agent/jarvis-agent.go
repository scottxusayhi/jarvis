package main

import (
	"git.oschina.net/k2ops/jarvis/agent/core"
	"git.oschina.net/k2ops/jarvis/agent/options"
	"git.oschina.net/k2ops/jarvis/agent/plugins"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
    "time"
)

var (
    connect chan bool
    id chan bool
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
	// options
	options.LoadCli()
	// logger
	initLogger()
	// connect
	core.KeepConnected(connect, id)
	go plugins.HandleMsg(connect)
    go report(id)
}
