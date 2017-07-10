//
// An echo service that receive messages from client and send back exactly the same one
//

package main

import (
	"net"
	"fmt"
	"os"
	"io"
	"git.oschina.net/k2ops/jarvis/utils"
	log "github.com/sirupsen/logrus"
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/server/handlers"
)

func HandleConnection(conn net.Conn) {
	log.WithFields(log.Fields{
		"localAddr": conn.LocalAddr(),
		"remoteAddr": conn.RemoteAddr(),
	}).Info("Client Connected")
	defer conn.Close()

	// send welcome message
	conn.Write(protocol.NewWelcomeMessage(conn.RemoteAddr().String(), conn.LocalAddr().String()).Serialize())

	//h := handlers.EchoHandler{}
	h2 := handlers.JarvisHandler{}
	content := make([]byte, 100)
	for {
		// clean input array for new request
		for c := range content {
			content[c] = 0
		}
		// read request
		n, err := conn.Read(content)
		if err == io.EOF {
			log.WithFields(log.Fields{
				"remoteAddr": conn.RemoteAddr(),
			}).Info("Connection Closed")
			break
		} else if err != nil {
			log.Error(err.Error())
		} else {
			//h.Handle(content[:n], conn)
			h2.Handle(content[:n], conn)
		}
	}
}

func main() {
	utils.InitLogger(log.DebugLevel)

	listener, err := net.Listen("tcp", ":2999")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Infof("Server started in pid %v", os.Getpid())

	// Listen for incoming connection
	for {
		conn, err := listener.Accept()
		// a timeout error
		// if err, ok := err.(*net.OpError); ok && err.Timeout() {}
		if err != nil {
			fmt.Println(err.Error())
		}
		go HandleConnection(conn)
	}
}
