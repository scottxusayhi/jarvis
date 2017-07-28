package api

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"net"
	"git.oschina.net/k2ops/jarvis/server/api/resource"
)

func StartServer(l net.Listener) error {
	defer log.Error("HTTP server failed to start or stopped")
	m := mux.NewRouter()
	// router, order matters
	m.HandleFunc("/api/v1", resource.RootHandler)
	m.HandleFunc("/api/v1/health", resource.HealthHandler)
	m.HandleFunc("/api/v1/hosts", resource.HostHandler)
	// serve everything in dir ./web under the path /
	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web"))))
	log.Info("HTTP REST server started")
	return http.Serve(l, m)
}
