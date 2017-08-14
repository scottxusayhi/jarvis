package tcp

import (
	log "github.com/sirupsen/logrus"
	"net"
)

func StartServer(l net.Listener) {
	defer log.Error("TCP server failed")
	log.Info("TCP server started")
	go OnlineCheck()
	for {
		log.Info("begin accept")
		conn, err := l.Accept()
		log.Info("accept returns")
		// a timeout error
		// if err, ok := err.(*net.OpError); ok && err.Timeout() {}
		if err != nil {
			log.Error(err.Error())
		}
		h := NewJarvisHandler(conn)
		h.Start()
	}
}
