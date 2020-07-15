package resource

import (
	"github.com/scottxusayhi/jarvis/server/api/helper"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"github.com/scottxusayhi/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
	"net/http"
	"database/sql"
)


func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	// common part
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	defer r.Body.Close()

	// CRUD
	switch r.Method {
	case http.MethodGet:
		getConfig(w, r)
		break
	case http.MethodPut:
		updateConfig(w, r)
		break
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	// search database
	b, err := mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}

	emails, err := b.GetEmailAlarmRecipients()
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	log.Debug(emails, err)

	// wrap api response and return
	config := model.Config{
		EmailAlarmRecipients: emails,
	}
	response, err := helper.WrapResponse(helper.SafeMarshalJsonObj(config), 0, "")
	if err != nil {
		log.Error(err.Error())
		helper.Write500Error(w, err.Error())
		return
	}
	w.Write(response)
}

func updateConfig(w http.ResponseWriter, r *http.Request) {
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