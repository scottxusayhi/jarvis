//
// An echo service that receive messages from client and send back exactly the same one
//

package main

import (
	"net"
	"fmt"
	"os"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"git.oschina.net/k2ops/jarvis/server/api"
	"git.oschina.net/k2ops/jarvis/server/tcp"
)


func main() {
	utils.InitLogger(log.DebugLevel)

	listener, err := net.Listen("tcp", ":2999")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Infof("Server started in pid %v", os.Getpid())

	// port multiplexing powered by github.com/soheilhy/cmux
	m := cmux.New(listener)
	httpL := m.Match(cmux.HTTP1Fast())
	tcpL := m.Match(cmux.Any())

	go api.NewServer(httpL)
	go tcp.NewServer(tcpL)

	m.Serve()

}
