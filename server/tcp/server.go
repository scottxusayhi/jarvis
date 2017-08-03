package tcp

import (
	"git.oschina.net/k2ops/jarvis/protocol"
	"git.oschina.net/k2ops/jarvis/server/tcp/handlers"
	"io"
	log "github.com/sirupsen/logrus"
	"net"
	"bufio"
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
	reader := bufio.NewReader(conn)
	for {
		// read request
		raw, err := reader.ReadBytes(protocol.Footer)
		if err == io.EOF {
			log.WithFields(log.Fields{
				"remoteAddr": conn.RemoteAddr(),
			}).Info("Connection Closed")
			break
		} else if err != nil {
			log.Error(err.Error())
		} else {
			h2.Handle(raw, conn)
		}
	}
}

func StartServer(l net.Listener) {
	defer log.Error("TCP server failed")
	log.Info("TCP server started")
	for {
		log.Info("begin accept")
		conn, err := l.Accept()
		log.Info("accept returns")
		// a timeout error
		// if err, ok := err.(*net.OpError); ok && err.Timeout() {}
		if err != nil {
			log.Error(err.Error())
		}
		go HandleConnection(conn)
	}
}