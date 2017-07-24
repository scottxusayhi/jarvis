package helper

import (
	"net/http"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func Write400Error (w http.ResponseWriter, message string) {
	response := model.ApiResBody{}
	response.Code = 1
	response.Message = message
	bytes, err := json.Marshal(model.ApiResBody{
		Code: 1,
		Message: message,
	})
	if err != nil {
		log.Error(err.Error())
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(bytes)
}
