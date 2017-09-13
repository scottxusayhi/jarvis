package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"git.oschina.net/k2ops/jarvis/server/backend"
	log "github.com/sirupsen/logrus"
)

func (m *JarvisMysqlBackend) SearchUser(query backend.Query) (users []model.User, err error) {
	db := m.db
	log.WithFields(log.Fields{
		"sql": "select * from jarvis.alarmlogs"+ query.SqlStringWhere(),
	}).Info("search alarms")
	rows, err := db.Query("SELECT * FROM jarvis.users" + query.SqlStringWhere())
	if err != nil {
		return
	}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.ReceiveAlarms,
		)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (m *JarvisMysqlBackend) GetEmailAlarmRecipients() (emails []string, err error) {
	query := backend.Query{
		"receiveAlarms": "1",
	}
	users, err := m.SearchUser(query)
	if err != nil {
		return
	}
	for _, user := range users {
		emails = append(emails, user.Email)
	}
	return
}