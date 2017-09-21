package utils

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger(level log.Level) {
	myTextFormatter := new(log.TextFormatter)
	myTextFormatter.FullTimestamp = true
       //myTextFormatter.TimestampFormat = log.defaultTimestampFormat

	log.SetFormatter(myTextFormatter)
	log.SetLevel(level)
}

func LogMsgSent(msg []byte, agent interface{}) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v -> sent message", agent)
}

func LogMsgReceived(msg []byte, agent interface{}) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v <- received message", agent)
}
