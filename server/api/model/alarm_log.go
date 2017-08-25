package model

import (
	"time"
)

type AlarmLog struct {
	Id int
	Time time.Time
	Alarm string
	Target string
	Value interface{}
	Status string
	Active bool
	Notified bool
}

func NewAlarmLog(alarm string, target string, value interface{}) AlarmLog {
	return AlarmLog{
		Time: time.Now(),
		Alarm:  alarm,
		Target: target,
		Value:  value,
	}
}

