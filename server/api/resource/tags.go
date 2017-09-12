package resource

import (
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"git.oschina.net/k2ops/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"encoding/json"
)


func TagsHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodPost:
		attachTag(w, r)
		break
	case http.MethodGet:
		getTags(w, r)
		break
	case http.MethodDelete:
		removeTag(w, r)
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func attachTag(w http.ResponseWriter, r *http.Request) {
	// get host id from url
	var systemId = systemId(r)
	// parse input
	newTags := make([]string, 10)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTags)
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
	// attach tag
	for _, tag := range newTags {
		err = backend.AttachTag(systemId, tag)
		if err != nil {
			log.Error(err.Error())
			helper.Write500Error(w, err.Error())
			return
		}
	}
	// return new tag list
	modifiedTags, err := backend.GetHostTags(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}
	w.Write(modifiedTags)
}

func getTags(w http.ResponseWriter, r *http.Request) {
	var systemId = systemId(r)
	// search database
	b, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	tags, err := b.GetHostTags(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(tags)
}

func removeTag(w http.ResponseWriter, r *http.Request) {
	// get host id from url
	var systemId = systemId(r)
	// parse input
	newTags := make([]string, 10)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTags)
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
	// remove tag
	for _, tag := range newTags {
		err = backend.RemoveTag(systemId, tag)
		if err != nil {
			log.Error(err.Error())
			helper.Write500Error(w, err.Error())
			return
		}
	}
	// return new tag list
	modifiedTags, err := backend.GetHostTags(systemId)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}
	w.Write(modifiedTags)
}


func listAllTags(w http.ResponseWriter, r *http.Request) {
	var err error
	backend, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	// query backend
	tags, err := backend.ListTags()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}

	bytes, err := json.Marshal(tags)
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
	}
	w.Write(bytes)

}