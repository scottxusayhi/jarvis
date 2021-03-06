package resource

import (
	"fmt"
	"github.com/scottxusayhi/jarvis/server/api/helper"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"github.com/scottxusayhi/jarvis/server/backend"
	"github.com/scottxusayhi/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func HostsHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
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
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
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
	query := backend.FromURLQuery(r.URL.Query())
	hosts, pageInfo, err := b.SearchHost(query)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	// wrap api response and return
	totalCount, err := b.CountHost(query)
	pageInfo.Size = len(hosts)
	pageInfo.TotalSize = totalCount
	pageInfo.CalcTotalPage()
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
	// determine delete type
	query := backend.FromURLQuery(r.URL.Query())
	delType, ok := query["type"]
	if !ok {
		helper.Write400Error(w, "query parameter type is required")
		return
	}
	if len(query) < 2 {
		helper.Write400Error(w, "at least one field is required (as query parameter)")
		return
	}

	// do delete
	b, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	switch delType {
	case "registry":
		affected, err := b.DeleteHostRegistry(query)
		if err != nil {
			log.Error(err.Error())
			helper.Write500Error(w, err.Error())
			return
		}
		helper.WriteResponse(w, http.StatusOK, 0, fmt.Sprintf("%v host(s) are un-registered", affected))
		break
	case "connection":
		b.DeleteHostConnection(query)
		break
	case "all":
		affected, err := b.DeleteHost(query)
		if err != nil {
			log.Error(err.Error())
			helper.Write500Error(w, err.Error())
		}
		helper.WriteResponse(w, http.StatusOK, 0, fmt.Sprintf("%v host(s) are totaly deleted", affected))
		break
	default:
		helper.Write500Error(w, "type must be one of [registry connection all]")
	}
}
