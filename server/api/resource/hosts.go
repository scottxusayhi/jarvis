package resource

import (
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"git.oschina.net/k2ops/jarvis/server/api/backend/mysql"
)

func HostHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	// CRUD
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

func registerHost(w http.ResponseWriter, r *http.Request) {
	host, err := model.ParseHost(r.Body)
	if err != nil {
		log.Error(err.Error())
		helper.Write400Error(w, err.Error())
		return
	}
	mysql.GetBackend().CreateHost(host)
	saved := mysql.GetBackend().GetOneHost(host.DataCenter, host.Rack, host.Slot, host.Hostname)
	w.Write(saved.JsonBytes())
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
