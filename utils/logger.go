package utils

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger(level log.Level) {
	myTextFormatter := new (log.TextFormatter)
	myTextFormatter.FullTimestamp = true
	myTextFormatter.TimestampFormat = log.DefaultTimestampFormat

	log.SetFormatter(myTextFormatter)
	log.SetLevel(level)
}


func LogMsgSent(msg []byte, agent string) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v -> sent message", agent)
}

func LogMsgReceived(msg []byte, agent string) {
	log.WithFields(log.Fields{
		"msg": string(msg),
	}).Infof("%v <- received message", agent)
}

