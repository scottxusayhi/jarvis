package api

import (
	"github.com/scottxusayhi/jarvis/server/api/resource"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func StartServer(l net.Listener) error {
	defer log.Error("HTTP server failed to start or stopped")
	m := mux.NewRouter()
	// router, order matters
	//m.HandleFunc("/api/v1", resource.RootHandler)
	//m.HandleFunc("/api/v1/health", resource.HealthHandler)
	//m.HandleFunc("/api/v1/list/{item}", resource.ListHandler)
	//m.HandleFunc("/api/v1/list", resource.AllListsHandler)
	//m.HandleFunc("/api/v1/hosts/{id}/tags", resource.TagsHandler)
	//m.HandleFunc("/api/v1/hosts/{id}", resource.OneHostHandler)
	//m.HandleFunc("/api/v1/hosts", resource.HostsHandler)
	//m.HandleFunc("/api/v1/config", resource.ConfigHandler)
	//// serve everything in dir ./web under the path /
	//m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/build"))))

	// new way
	apiRoute := m.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/", resource.RootHandler)
	apiRoute.HandleFunc("/health", resource.HealthHandler)
	apiRoute.HandleFunc("/list/{item}", resource.ListHandler)
	apiRoute.HandleFunc("/list", resource.AllListsHandler)
	apiRoute.HandleFunc("/hosts/{id}/tags", resource.TagsHandler)
	apiRoute.HandleFunc("/hosts/{id}", resource.OneHostHandler)
	apiRoute.HandleFunc("/hosts", resource.HostsHandler)
	apiRoute.HandleFunc("/config", resource.ConfigHandler)

	m.PathPrefix("/css").Handler(http.FileServer(http.Dir("./web/build/")))
	m.PathPrefix("/fonts").Handler(http.FileServer(http.Dir("./web/build/")))
	m.PathPrefix("/img").Handler(http.FileServer(http.Dir("./web/build/")))
	m.PathPrefix("/static").Handler(http.FileServer(http.Dir("./web/build/")))
	m.PathPrefix("/").HandlerFunc(IndexHandler("./web/build/index.html"))
	//m.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/build"))))
	log.Info("HTTP REST server started")
	return http.Serve(l, m)
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !isLoggedIn(r) {
			http.Redirect(w, r, )
		}
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}


