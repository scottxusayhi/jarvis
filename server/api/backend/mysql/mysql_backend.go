package mysql

import (
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"git.oschina.net/k2ops/jarvis/server/api/backend"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"git.oschina.net/k2ops/jarvis/utils"
	"fmt"
)

var (
	b *JarvisMysqlBackend = nil
)

type JarvisMysqlBackend struct {
	host string
	port uint16
	db *sql.DB
	stmtGetOneHost *sql.Stmt
	stmtInsertHost *sql.Stmt
}

func (m *JarvisMysqlBackend) prepareStatements() error {
	db := m.db
	var err error
	m.stmtGetOneHost, err = db.Prepare("select * from hosts where datacenter=? and rack=? and slot=?;")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	
	m.stmtInsertHost, err = db.Prepare("insert into hosts(datacenter, rack, slot, tags, osExpected, osDetected, cpuExpected, cpuDetected, memExpected, memDetected, diskExpected, diskDetected, networkExpected, networkDetected, registered) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, TRUE);")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (m *JarvisMysqlBackend) CreateHost(h model.Host) error {
	log.WithFields(log.Fields{
		"sql": m.stmtInsertHost,
	}).Info("create host")
	result, err := m.stmtInsertHost.Exec(
		h.DataCenter,
		h.Rack,
		h.Slot,
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

func (m *JarvisMysqlBackend) CountHost(q backend.Query) (int, error) {
	db := m.db
	var count int
	log.WithFields(log.Fields{
		"sql": "select count(*) from hosts " + q.SqlString(),
	}).Info("count hosts")
	err := db.QueryRow("select count(*) from hosts " + q.SqlString()).Scan(&count)
	return count, err
}

func (m *JarvisMysqlBackend) SearchHost(q backend.Query) ([]model.Host, error) {
	var hosts []model.Host
	db := m.db
	log.WithFields(log.Fields{
		"sql": "select * from hosts " + q.SqlString(),
	}).Info("search hosts")
	rows, err := db.Query("select * from hosts " + q.SqlString())
	if err != nil {
		log.Error("mysql error: " + err.Error())
		return nil, err
	}
	for rows.Next() {
		host := model.Host{}
		err := rows.Scan(
			&host.SystemId,
			&host.DataCenter,
			&host.Rack,
			&host.Slot,
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

func (m *JarvisMysqlBackend) GetOneHost(dc string, rack string, slot string) (*model.Host, error) {
	query := backend.Query{
		"datacenter": dc,
		"rack": rack,
		"slot": slot,
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


// delete both registry and connection info,
// just delete the db record
// need restart agent afterwards
func (m *JarvisMysqlBackend) DeleteHost(q backend.Query) (int64, error) {
	db := m.db
	stmt := "DELETE FROM hosts " + q.SqlString()
	log.WithFields(log.Fields{
		"sql": stmt,
	}).Info("delete host all info")
	result, err := db.Exec(stmt)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return result.RowsAffected()
}


// only delete host registry info
func (m *JarvisMysqlBackend) DeleteHostRegistry (q backend.Query) (int64, error) {
	randDatacenter := utils.UnknownDataCenter()
	randRack := utils.UnknownRack()
	randSlot := utils.UnknownSlot()
	db := m.db
	stmt := fmt.Sprintf(`UPDATE hosts SET
	datacenter="%v",
	rack="%v",
	slot="%v",
	owner=DEFAULT,
	osExpected="{}",
	cpuExpected="{}",
	memExpected="{}",
	diskExpected="[]",
	networkExpected="{}",
	registered=0,
	matched=0,
	createdAt="0001-01-01 00:00:00",
	updatedAt="0001-01-01 00:00:00" %v`, randDatacenter, randRack, randSlot, q.SqlString())
	log.WithFields(log.Fields{
		"sql": stmt,
	}).Info("clean host registry info")
	result, err := db.Exec(stmt)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return result.RowsAffected()
}

// only delete host connection info
// need restart agent afterwards
func (m *JarvisMysqlBackend) DeleteHostConnection (q backend.Query) (int64, error) {
	db := m.db
	stmt := fmt.Sprintf(`UPDATE hosts SET
	osDetected="{}",
	cpuDetected="{}",
	memDetected="{}",
	diskDetected="[]",
	networkDetected="{}",
	connected=0,
	matched=0,
	online=0,
	firstSeenAt="0001-01-01 00:00:00",
	lastSeenAt="0001-01-01 00:00:00" %v`, q.SqlString())
	log.WithFields(log.Fields{
		"sql": stmt,
	}).Info("clean host connection info")
	result, err := db.Exec(stmt)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}
	return result.RowsAffected()
}


func GetBackend () (*JarvisMysqlBackend, error) {
	if b == nil {
		dsn := "root:passw0rd@tcp(localhost:3306)/jarvis?parseTime=true"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Error(err.Error())
		}
		err = db.Ping()
		if err != nil {
			log.Error(err.Error())
		}
		log.WithFields(log.Fields{
			"dsn": dsn,
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