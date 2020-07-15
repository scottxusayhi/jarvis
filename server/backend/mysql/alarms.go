package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"github.com/scottxusayhi/jarvis/server/backend"
	log "github.com/sirupsen/logrus"
)

func (m *JarvisMysqlBackend) SearchAlarms(query backend.Query) (als []model.AlarmLog, err error) {
	db := m.db
	log.WithFields(log.Fields{
		"sql": "select * from jarvis.alarmlogs"+ query.SqlStringWhere(),
	}).Info("search alarms")
	rows, err := db.Query("SELECT * FROM jarvis.alarmlogs" + query.SqlStringWhere())
	if err != nil {
		return
	}
	for rows.Next() {
		alarm := model.AlarmLog{}
		err = rows.Scan(
			&alarm.Id,
			&alarm.Time,
			&alarm.Alarm,
			&alarm.Target,
			&alarm.Value,
			&alarm.Status,
			&alarm.Active,
			&alarm.Notified,
		)
		if err != nil {
			return
		}
		als = append(als, alarm)
	}
	return
}


func (m *JarvisMysqlBackend) GetActiveAlarms() (als []model.AlarmLog, err error) {
	return m.SearchAlarms(backend.Query{"active": "1"})
}


func (m *JarvisMysqlBackend) AlarmExistsAndActive(e model.AlarmLog) (active bool, err error) {
	als, err := m.SearchAlarms(backend.Query{
		"alarm": e.Alarm,
		"target": e.Target,
		"active": "1",
	})
	if err != nil {
		return
	}
	if len(als)==0 {
		active = false
	} else {
		active = true
	}
	return
}

func (m *JarvisMysqlBackend) ClearAlarm(e model.AlarmLog) (n int, err error) {
	db := m.db
	result, err := db.Exec("UPDATE jarvis.alarmlogs SET active=FALSE WHERE alarm=? and target=? AND active=TRUE",
		e.Alarm,
		e.Target,
	)
	if err != nil {
		return
	}
	affected, err := result.RowsAffected()
	n = int(affected)
	return
}

func (m *JarvisMysqlBackend) GetOfflineHosts() (hosts []model.Host, err error) {
	db := m.db
	rows, err := db.Query("SELECT systemId, online FROM jarvis.hosts WHERE online=FALSE")
	if err != nil {
		return
	}
	for rows.Next() {
		host := model.Host{}
		err = rows.Scan(
			&host.SystemId,
			&host.Online,
		)
		if err != nil {
			return
		}
		hosts = append(hosts, host)
	}
	return
}


func (m *JarvisMysqlBackend) WriteAlarmLog(e *model.AlarmLog) (err error) {
	db := m.db
	result, err := db.Exec("INSERT INTO jarvis.alarmlogs(time, alarm, target, value, status, active, notified) VALUES(?, ?, ?, ?, ?, ?, ?)",
		e.Time,
		e.Alarm,
		e.Target,
		e.Value,
		e.Status,
		e.Active,
		e.Notified,
	)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	e.Id = int(id)
	e.Active = false
	e.Notified = false
	return
}