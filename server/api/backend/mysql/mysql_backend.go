package mysql

import (
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"git.oschina.net/k2ops/jarvis/server/api/backend"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
)

var (
	b *JarvisMysqlBackend = nil
)

type JarvisMysqlBackend struct {
	host string
	port uint16
	db *sql.DB
	stmtSelectHosts *sql.Stmt
	stmtGetOneHost *sql.Stmt
	stmtInsertHost *sql.Stmt
}

func (m *JarvisMysqlBackend) prepareStatements() error {
	db := m.db
	var err error
	m.stmtGetOneHost, err = db.Prepare("select * from hosts where datacenter=? and rack=? and slot=? and hostname=?;")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	m.stmtSelectHosts, err = db.Prepare("select * from hosts WHERE ?;")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	
	m.stmtInsertHost, err = db.Prepare("insert into hosts(datacenter, rack, slot, hostname, tags, osExpected, osDetected, cpuExpected, cpuDetected, memExpected, memDetected, diskExpected, diskDetected, networkExpected, networkDetected) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (m *JarvisMysqlBackend) CreateHost(h model.Host) error {
	result, err := m.stmtInsertHost.Exec(
		h.DataCenter,
		h.Rack,
		h.Slot,
		h.Hostname,
		helper.SafeMarshalJsonArray(h.Tags),
		helper.SafeMarshalJsonObj(h.OsExpected),
		"{}",
		helper.SafeMarshalJsonObj(h.CpuExpected),
		"{}",
		helper.SafeMarshalJsonObj(h.MemExpected),
		"{}",
		helper.SafeMarshalJsonArray(h.DiskExpected),
		"[]",
		helper.SafeMarshalJsonObj(h.NetworkExpected),
		"{}",
	)
	if err != nil {
		log.Error("mysql error insert: " + err.Error())
		return err
	}
	id, _ := result.LastInsertId()
	log.WithFields(log.Fields{
		"insertId": id,
	}).Info("host created")
	return nil
}

func (m *JarvisMysqlBackend) SearchHost(q backend.Query) ([]model.Host, error) {
	var hosts []model.Host
	db := m.db
	rows, err := db.Query("select * from hosts WHERE " + q.String())
	if err != nil {
		log.Error("mysql error: " + err.Error())
		return nil, err
	}
	for rows.Next() {
		host := model.Host{}
		err := rows.Scan(
			&host.DataCenter,
			&host.Rack,
			&host.Slot,
			&host.Hostname,
			&host.Tags,
			&host.Owner,
			&host.OsExpected,
			&host.OsDetected,
			&host.CpuExpected,
			&host.CpuDetected,
			&host.MemExpected,
			&host.MemDetected,
			&host.DiskExpected,
			&host.DiskDetected,
			&host.NetworkExpected,
			&host.NetworkDetected,
			&host.Registered,
			&host.Connected,
			&host.Matched,
			&host.Online,
			&host.HealthStatus,
			&host.FirstSeenAt,
			&host.LastSeenAt,
			&host.CreatedAt,
			&host.UpdatedAt,
		)
		if err != nil {
			log.Error(err.Error())
		}
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func (m *JarvisMysqlBackend) GetOneHost(dc string, rack string, slot string, hostname string) (*model.Host, error) {
	query := backend.Query{
		"datacenter": dc,
		"rack": rack,
		"slot": slot,
		"hostname": hostname,
	}

	hosts, err := m.SearchHost(query)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(hosts)>0 {
		return &hosts[0], nil
	}
	return nil, sql.ErrNoRows
}

func (m *JarvisMysqlBackend) UpdateHost(q backend.Query, h model.Host) error {
	panic("implement me")
}

func (m *JarvisMysqlBackend) DeleteHost(q backend.Query) error {
	panic("implement me")
}


func GetBackend () (*JarvisMysqlBackend, error) {
	if b == nil {
		db, err := sql.Open("mysql", "root:passw0rd@tcp(localhost:3306)/jarvis?parseTime=true")
		if err != nil {
			log.Error(err.Error())
		}
		err = db.Ping()
		if err != nil {
			log.Error(err.Error())
		}
		log.WithFields(log.Fields{
			"addr": "default",
		}).Info("MySQL connected.")
		b = &JarvisMysqlBackend{
			host: "localhost",
			port: 2379,
			db: db,
		}
		b.prepareStatements()
	}
	return b, nil
}