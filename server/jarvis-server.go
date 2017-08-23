//
// An echo service that receive messages from client and send back exactly the same one
//

package main

import (
	"git.oschina.net/k2ops/jarvis/server/api"
	"git.oschina.net/k2ops/jarvis/server/tcp"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"net"
	"os"
	"git.oschina.net/k2ops/jarvis/server/options"
)

func initLogger() {
	utils.InitLogger(log.InfoLevel)
	if options.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug enabled.")
	}
}

func main() {
	options.LoadCli()
	initLogger()
	both()
}

func both() {
	// open port
	listener, err := net.Listen("tcp", ":2999")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithFields(log.Fields{
		"pid":  os.Getpid(),
		"port": ":2999",
	}).Info("Server started.")

	// port multiplexing powered by github.com/soheilhy/cmux
	m := cmux.New(listener)
	httpL := m.Match(cmux.HTTP1Fast())
	tcpL := m.Match(cmux.Any())

	// serve http including rest api and web frontend
	go api.StartServer(httpL)
	// serve tcp communication between master and slaves
	go tcp.StartServer(tcpL)

	m.Serve()
}

func onlyTcp() {
	// open port
	listener, err := net.Listen("tcp", ":2999")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.WithFields(log.Fields{
		"pid":  os.Getpid(),
		"port": ":2999",
	}).Info("Server started.")

	// serve tcp communication between master and slaves
	tcp.StartServer(listener)
}
