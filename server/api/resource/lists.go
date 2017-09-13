package resource

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"encoding/json"
	"git.oschina.net/k2ops/jarvis/server/backend/mysql"
)

func item(r *http.Request) (string) {
	return mux.Vars(r)["item"]
}


func ListHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodGet:
		listItem(w, r)
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listItem(w http.ResponseWriter, r *http.Request) {
	item := item(r)
	switch item {
	case "tags":
		listAllTags(w, r)
		break
	default:
		log.Error("unknow item to list: " + item)
		w.WriteHeader(http.StatusNotFound)
	}
}

func AllListsHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodGet:
		listAllItems(w, r)
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func listAllItems(w http.ResponseWriter, r *http.Request) {
	var err error
	backend, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	// tags
	tags, err := backend.ListTags()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}

	// datacenters
	dcs, err := backend.ListDatacenters()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}

	// racks
	racks, err := backend.ListRacks()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}

	// owners
	owners, err := backend.ListOwner()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}

	// compose result
	result := make(map[string]interface{})
	result["tags"] = tags
	result["datacenters"] = dcs
	result["racks"] = racks
	result["owners"] = owners

	bytes, err := json.Marshal(result)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}
	w.Write(bytes)
}
