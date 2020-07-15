package resource

import (
	"fmt"
	"github.com/scottxusayhi/jarvis/server/api/helper"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"github.com/scottxusayhi/jarvis/server/backend"
	"github.com/scottxusayhi/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
)

func systemId(r *http.Request) (string) {
	return mux.Vars(r)["id"]
}

func OneHostHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodPost:
		postRegHost(w, r)
		break
	case http.MethodGet:
		getOneHost(w, r)
		break
	case http.MethodPut:
		updateOneHost(w, r)
		break
	//case http.MethodDelete:
	//	deleteOneHost(w, r)
	//	break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func postRegHost(w http.ResponseWriter, r *http.Request) {
	// get host id from url
	var systemId = systemId(r)
	// parse input
	inputHost, err := model.ParseHost(r.Body)
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
	// check host
	host, err := backend.GetOneHostById(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write400Error(w, err.Error())
		return
	}
	if host.Registered {
		helper.Write400Error(w, fmt.Sprintf("host systemId=%v is already registered", host.SystemId))
		return
	}

	// post reg host
	err = backend.PostRegHost(inputHost, systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write400Error(w, err.Error())
	}

	// load and return
	saved, err := backend.GetOneHostById(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(saved.JsonBytes())
}

func getOneHost(w http.ResponseWriter, r *http.Request) {
	var systemId = systemId(r)
	// search database
	b, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	host, err := b.GetOneHostById(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	// wrap api response and return
	response, err := helper.WrapResponse(host.JsonBytes(), 0, "")
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(response)
}

func updateOneHost(w http.ResponseWriter, r *http.Request) {
	var systemId = systemId(r)
	var err error
	backend, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	// parse updates
	update, err := model.ParseUpdatableFields(r.Body)
	if err != nil {
		log.Error(err.Error())
		helper.Write400Error(w, err.Error())
		return
	}
	// do update
	err = backend.UpdateHostById(systemId, update)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	// load and return
	host, err := backend.GetOneHostById(systemId)
	if err != nil {
		log.Error(err.Error())
		if err==sql.ErrNoRows {
			helper.Write400Error(w, err.Error())
		} else {
			helper.Write500Error(w, err.Error())
		}
		return
	}

	result, err := helper.WrapResponseSuccess(host.JsonBytes())
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(result)
}

func deleteOneHost(w http.ResponseWriter, r *http.Request) {
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
