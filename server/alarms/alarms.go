package alarms

import (
	"time"
	"github.com/scottxusayhi/jarvis/server/backend/mysql"
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"strings"
	"strconv"
)

const (
	STATUS_UNDEFINED="UNDEFINED"
	STATUS_CLEAR="CLEAR"
	STATUS_CRITICAL="CRITICAL"
)

var (
	AllAlarms = map[string]AlarmEvaluator {
		"host_alarm": hostAlarm{},
	}
	be *mysql.JarvisMysqlBackend = nil
)

func Start() {
	var err error
	be, err = mysql.GetBackend()
	if err != nil {
		log.Error(err.Error())
		return
	}
	for k, v := range AllAlarms {
		go v.Check(be)
		log.WithFields(log.Fields{
			"name": k,
		}).Info("alarm started")
	}
}

type AlarmEvaluator interface {
	Name() string
	Desc() string
	CheckEvery() time.Duration
	Check(backend *mysql.JarvisMysqlBackend)
	Evaluate(value interface{}) string
	Handle(e model.AlarmLog)
	Notify(e *model.AlarmLog)
}


type hostAlarm struct {}

func (alarm hostAlarm) Notify(e *model.AlarmLog) {
	var err error
	to, err := be.GetEmailAlarmRecipients()
	if err != nil {
		log.WithError(err).Error("notify alarm failed: failed to get email recipients")
		return
	}
	switch e.Status {
	case STATUS_CRITICAL: {
		subj := fmt.Sprintf("[From Jarvis] Host Offline")
		body := fmt.Sprintf("systemId = %v", e.Target)
		err = SendMail(to, subj, body)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.WithFields(log.Fields{
			"host_id": e.Target,
		}).Debug("host offline notified via email")
		e.Notified = true
		break
	}
	case STATUS_CLEAR:
		subj := fmt.Sprintf("[From Jarvis] Host Recovery (Online)")
		body := fmt.Sprintf("systemId = %v", e.Target)
		err = SendMail(to, subj, body)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.WithFields(log.Fields{
			"host_id": e.Target,
		}).Debug("host recovery (online) notified via email")
		e.Notified = true
		break
	default:
		log.Info("no need to notify: " + e.Status)
	}
	return
}

func (alarm hostAlarm) Handle(e model.AlarmLog) {
	log.WithFields(log.Fields{
		"alarm_log": e,
	}).Debug("============ handle alarm log start ============")
	e.Status = alarm.Evaluate(e.Value)
	log.WithField("status", e.Status).Debug()
	switch e.Status {
	case STATUS_CRITICAL:
		active, err := be.AlarmExistsAndActive(e)
		if err != nil {
			log.Error(err.Error())
			break
		}
		if !active {
			log.WithFields(log.Fields{
				"alarm_log": e,
			}).Debug("new alarm")
			e.Active = true
			alarm.Notify(&e)
		} else {
			log.WithField("alarm_log", e).Debug("this alarm is already active")
		}
		break
	case STATUS_CLEAR:
		n, err := be.ClearAlarm(e)
		if err != nil {
			log.Error(err.Error())
			break
		}
		if n>0 {
			alarm.Notify(&e)
		}
		break
	case STATUS_UNDEFINED:
		break
	default:
		log.Error("unknown alarm status: " + e.Status)
	}
	log.WithField("alarm_log", e).Debug("will save alarm log")
	be.WriteAlarmLog(&e)
}

func (alarm hostAlarm) Name() string {
	return "host_alarm"
}

func (alarm hostAlarm) Desc() string {
	return "host online status"
}

func (alarm hostAlarm) CheckEvery() time.Duration {
	return 30*time.Second
}

func (alarm hostAlarm) Check(backend *mysql.JarvisMysqlBackend) {
	for ; ;time.Sleep(alarm.CheckEvery())  {
		// currently active alarms
		activeAlarms, err := be.GetActiveAlarms()
		if err != nil {
			log.Error(err.Error())
			continue
		}
		log.WithFields(log.Fields{
			"count": len(activeAlarms),
		}).Info("active alarms")

		// currently offline hosts
		offlineHosts, err := be.GetOfflineHosts()
		if err != nil {
			log.Error(err.Error())
			continue
		}
		log.WithFields(log.Fields{
			"count": len(offlineHosts),
		}).Info("offline hosts")

		cleared := 0
		new := 0
		// see if wen can clear some alarms
		for _, a := range activeAlarms {
			if !inHostList(a.Target, offlineHosts) {
				alarm.Handle(model.NewAlarmLog(alarm.Name(), a.Target, true))
				cleared++
			}
		}
		// new alarms
		for _, h := range offlineHosts {
			if !inAlarmList(h.SystemId, activeAlarms) {
				alarm.Handle(model.NewAlarmLog(alarm.Name(), fmt.Sprintf("%v", h.SystemId), false))
				new++
			}
		}
		log.WithFields(log.Fields{
			"cleared_alarm(s)": cleared,
			"new_alarm(s)": new,
		}).Info("host_alarm checked")
	}
}

func (alarm hostAlarm) Evaluate(value interface{}) string {
	boolValue, ok := value.(bool)
	if !ok {
		log.WithFields(log.Fields{
			"expected": "bool",
			"got": fmt.Sprintf("%T", value),
		}).Error("evaluate host alarm value error")
		return STATUS_UNDEFINED
	}

	if boolValue {
		return STATUS_CLEAR
	} else {
		return STATUS_CRITICAL
	}
}

func inAlarmList(id int, list []model.AlarmLog) bool {
	for _, a := range list {
		if strings.EqualFold(a.Target, fmt.Sprintf("%v", id)) {
			return true
		}
	}
	return false
}

func inHostList(id string, list []model.Host) bool {
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return false
	}
	for _, a := range list {
		if a.SystemId==int(intId) {
			return true
		}
	}
	return false
}