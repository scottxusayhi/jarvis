package resource

import (
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"git.oschina.net/k2ops/jarvis/server/api/backend/mysql"
	"strings"
	//"git.oschina.net/k2ops/jarvis/server/api/backend"
	"git.oschina.net/k2ops/jarvis/server/api/backend"
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
	// parse input
	host, err := model.ParseHost(r.Body)
	if err != nil {
		log.Error(err.Error())
		helper.Write400Error(w, err.Error())
		return
	}
	// get db connection
	backend, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	// save to db
	err = backend.CreateHost(host)
	if err != nil {
		log.Error(err.Error())
		if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
			helper.Write400Error(w, err.Error())
			return
		}
		helper.Write500Error(w, err.Error())
		return
	}
	// load and return
	saved, err := backend.GetOneHost(host.DataCenter, host.Rack, host.Slot)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(saved.JsonBytes())
}

func searchHosts(w http.ResponseWriter, r *http.Request) {
	// search database
	b, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	hosts, err := b.SearchHost(backend.FromURLQuery(r.URL.Query()))
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	// wrap api response and return
	pageInfo := helper.DefaultPageInfo()
	response, err := helper.WrapHostListResponse(0, "", hosts, pageInfo)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(response)
}

func updateHost(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	w.Write([]byte("update host"))
}

func deleteHost(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Info(query)
	if len(query)==0 {
		helper.Write400Error(w, "query parameter is required")
		return
	}
}
