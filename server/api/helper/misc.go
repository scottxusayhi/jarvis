package helper

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func SafeMarshalJsonArray(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Error(err.Error())
		return []byte("[]")
	}
	return data
}

func SafeMarshalJsonObj(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Error(err.Error())
		return []byte("{}")
	}
	return data
}
