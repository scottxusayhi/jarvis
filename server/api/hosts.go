package api

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func hostHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerHost(w, r)
		break
	case http.MethodGet:
		searchHosts(w, r)
		break
	case http.MethodPut:
		updateHost(w, r)
		break
	case http.MethodDelete:
		deleteHost(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func registerHost (w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	w.Write([]byte("create host"))
}

func searchHosts(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	w.Write([]byte("search host"))
}

func updateHost(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	w.Write([]byte("update host"))
}

func deleteHost(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	w.Write([]byte("update host"))
}

