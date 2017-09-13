package api

import (
	"git.oschina.net/k2ops/jarvis/server/api/resource"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func StartServer(l net.Listener) error {
	defer log.Error("HTTP server failed to start or stopped")
	m := mux.NewRouter()
	// router, order matters
	m.HandleFunc("/api/v1", resource.RootHandler)
	m.HandleFunc("/api/v1/health", resource.HealthHandler)
	m.HandleFunc("/api/v1/list/{item}", resource.ListHandler)
	m.HandleFunc("/api/v1/list", resource.AllListsHandler)
	m.HandleFunc("/api/v1/hosts/{id}/tags", resource.TagsHandler)
	m.HandleFunc("/api/v1/hosts/{id}", resource.OneHostHandler)
	m.HandleFunc("/api/v1/hosts", resource.HostsHandler)
	m.HandleFunc("/api/v1/config", resource.ConfigHandler)
	// serve everything in dir ./web under the path /
	m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/build"))))
	log.Info("HTTP REST server started")
	return http.Serve(l, m)
}
