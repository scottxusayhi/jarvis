package model

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ReceiveAlarms bool `json:"receiveAlarms"`
}

type Config struct {
	EmailAlarmRecipients []string `json:"emailAlarmRecipients"`
}

func (config *Config) JsonBytes() []byte {
	bytes, err := json.Marshal(config)
	if err != nil {
		log.Error(err.Error())
		return []byte("{}")
	}
	return bytes
}

func (config *Config) JsonString() string {
	return string(config.JsonBytes())
}


