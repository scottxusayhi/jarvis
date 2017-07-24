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


