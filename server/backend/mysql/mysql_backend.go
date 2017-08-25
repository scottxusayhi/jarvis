package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"git.oschina.net/k2ops/jarvis/server/backend"
	"git.oschina.net/k2ops/jarvis/utils"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"time"
	"git.oschina.net/k2ops/jarvis/protocol"
	"sync"
	"git.oschina.net/k2ops/jarvis/server/options"
)

var (
	b *JarvisMysqlBackend = nil
	mysqlConnectionLock = sync.Mutex{}
)

type JarvisMysqlBackend struct {
	host           string
	port           uint16
	db             *sql.DB
	stmtGetOneHost *sql.Stmt
	stmtInsertHost *sql.Stmt
}

// for http api methods
func (m *JarvisMysqlBackend) prepareStatements() error {
	db := m.db
	var err error
	m.stmtGetOneHost, err = db.Prepare("select * from hosts where datacenter=? and rack=? and slot=?;")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	m.stmtInsertHost, err = db.Prepare("insert into hosts(systemId, datacenter, rack, slot, tags, osExpected, osDetected, cpuExpected, cpuDetected, memExpected, memDetected, diskExpected, diskDetected, networkExpected, networkDetected, registered) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, TRUE);")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (m *JarvisMysqlBackend) CreateHost(h model.Host) (err error) {
	// start a database transaction
	db := m.db
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// allocate new id
	result, err := tx.Exec("INSERT INTO jarvis.ids(status) VALUES ('used')")
	if err != nil {
		tx.Rollback()
		return err
	}
	newId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// register host
	log.Info("register host")
	result, err = tx.Exec("insert into hosts(systemId, datacenter, rack, slot, tags, osExpected, osDetected, cpuExpected, cpuDetected, memExpected, memDetected, diskExpected, diskDetected, networkExpected, networkDetected, registered) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, TRUE);",
		newId,
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
		tx.Rollback()
		return err
	}
	id, _ := result.LastInsertId()
	log.WithFields(log.Fields{
		"systemId": id,
	}).Info("host created")
	err = tx.Commit()
	if err != nil {
		return err
	}
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

func (m *JarvisMysqlBackend) SearchHost(q backend.Query) (hosts []model.Host, pageInfo helper.PageInfo, err error) {
	db := m.db
	pageInfo = backend.PageInfo(q)
	query := fmt.Sprintf("select * from jarvis.hosts %v %v", q.SqlString(), pageInfo.SqlString())
	log.WithFields(log.Fields{
		"sql": query,
	}).Info("search hosts")
	rows, err := db.Query(query)
	if err != nil {
		log.Error("mysql error: " + err.Error())
		return
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
	return
}

func (m *JarvisMysqlBackend) GetOneHost(dc string, rack string, slot string) (*model.Host, error) {
	query := backend.Query{
		"datacenter": dc,
		"rack":       rack,
		"slot":       slot,
	}

	hosts, _, err := m.SearchHost(query)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(hosts) > 0 {
		return &hosts[0], nil
	}
	return nil, sql.ErrNoRows
}

func (m *JarvisMysqlBackend) GetOneHostById(aid string) (*model.Host, error) {
	query := backend.Query{
		"systemId": aid,
	}
	hosts, _, err := m.SearchHost(query)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(hosts) > 0 {
		return &hosts[0], nil
	}
	return nil, sql.ErrNoRows
}


func (m *JarvisMysqlBackend) UpdateHost(q backend.Query, update map[string]interface{}) error {
	db := m.db
	updateSql := fmt.Sprintf("UPDATE jarvis.hosts SET %v %v", model.UpdateSqlString(update), q.SqlString())
	log.WithFields(log.Fields{
		"sql": updateSql,
		"values": model.UpdateValues(update),
	}).Info("update host")
	result, err := db.Exec(updateSql, model.UpdateValues(update)...)
	if err != nil {
		return err
	}
	affectedRows, err := result.RowsAffected()
	log.WithFields(log.Fields{
		"rowsAffected": affectedRows,
	}).Info("updated host")
	return err
}

func (m *JarvisMysqlBackend) UpdateHostById(hostId string, update map[string]interface{}) error {
	query := backend.Query{
		"systemId": hostId,
	}
	return m.UpdateHost(query, update)
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
func (m *JarvisMysqlBackend) DeleteHostRegistry(q backend.Query) (int64, error) {
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
func (m *JarvisMysqlBackend) DeleteHostConnection(q backend.Query) (int64, error) {
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

// for tcp protocol
func (m *JarvisMysqlBackend) PreserveId() (int64, error) {
	db := m.db
	result, err := db.Exec("INSERT INTO jarvis.ids(status) VALUES ('preserved')")
	if err != nil {
		return 0, err
	}
	newId, err := result.LastInsertId()
	log.WithFields(log.Fields{
		"agentId": newId,
	}).Info("New agent id preserved")
	return newId, err
}

func (m *JarvisMysqlBackend) UpdateConnectionInfo(agentId int64) error {
	var count int
	db := m.db
	db.QueryRow("SELECT count(*) from jarvis.hosts WHERE systemId=?", agentId).Scan(&count)
	if count == 0 {
		log.WithFields(log.Fields{
			"agentId": agentId,
		}).Info("No connection info found, going to create")
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		_, err = tx.Exec("insert into jarvis.hosts(systemId, datacenter, rack, slot, tags, osExpected, osDetected, cpuExpected, cpuDetected, memExpected, memDetected, diskExpected, diskDetected, networkExpected, networkDetected, connected, online) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			agentId,
			utils.UnknownDataCenter(),
			utils.UnknownRack(),
			utils.UnknownSlot(),
			"[]",
			"{}",
			"{}",
			"{}",
			"{}",
			"{}",
			"{}",
			"[]",
			"[]",
			"{}",
			"{}",
			true,
			true,
		)
		if err != nil {
			return err
		}
		result, err := tx.Exec("update jarvis.ids set status=? where nextId=?", "used", agentId)
		if err != nil {
			tx.Rollback()
			return err
		}
		updated, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return err
		}
		if updated == 0 {
			tx.Rollback()
			return errors.New(fmt.Sprintf("agent id %v is missing in id table", agentId))
		}
		return tx.Commit()
	} else {
		log.WithFields(log.Fields{
			"agentId": agentId,
		}).Info("Update connection info")
		_, err := db.Exec("UPDATE jarvis.hosts SET online=TRUE WHERE systemId=?", agentId)
		return err
	}
}

func (m *JarvisMysqlBackend) UpdateHeartBeat(id int64, updateTime time.Time) error {
	db := m.db
	_, err := db.Exec("UPDATE jarvis.hosts SET online=TRUE, firstSeenAt=? WHERE systemId=? AND firstSeenAt=?", updateTime, id, "0001-01-01 00:00:00")
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE jarvis.hosts SET online=TRUE, lastSeenAt=? WHERE systemId=?", updateTime, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *JarvisMysqlBackend) UpdateHostConfig(id int64, osDetected protocol.OsInfo, cpuDetected protocol.CpuInfo, memDetected protocol.MemInfo, diskDetected protocol.HostDisks, networkDetected protocol.NetworkInfo, match bool) error {
	db := m.db
	_, err := db.Exec("UPDATE jarvis.hosts SET osDetected=?, cpuDetected=?, memDetected=?, diskDetected=?, networkDetected=?, matched=? WHERE systemId=?",
		helper.SafeMarshalJsonObj(osDetected),
		helper.SafeMarshalJsonObj(cpuDetected),
		helper.SafeMarshalJsonObj(memDetected),
		helper.SafeMarshalJsonArray(diskDetected),
		helper.SafeMarshalJsonObj(networkDetected),
		match,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *JarvisMysqlBackend) QueryExpectedConfig(aid string) (*model.OsInfo, *model.CpuInfo, *model.MemInfo, *model.HostDisks, *model.NetworkInfo, error) {
	host, err := m.GetOneHostById(aid)
	if err!=nil {
		return nil, nil, nil, nil, nil, err
	}
	return &host.OsExpected, &host.CpuExpected, &host.MemExpected, &host.DiskExpected, &host.NetworkExpected, nil
}

func (m *JarvisMysqlBackend) MarkOffline(id int64) error {
	db := m.db
	var count int
	db.QueryRow("SELECT count(*) FROM jarvis.hosts WHERE systemId=?", id).Scan(&count)
	if count == 0 {
		return errors.New(fmt.Sprintf("agent id %v not found", id))
	}
	_, err := db.Exec("UPDATE jarvis.hosts SET online=FALSE WHERE systemId=?", id)
	return err
}

func (m *JarvisMysqlBackend) MarkOnline(id int64) error {
	db := m.db
	var count int
	db.QueryRow("SELECT count(*) FROM jarvis.hosts WHERE systemId=?", id).Scan(&count)
	if count == 0 {
		return errors.New(fmt.Sprintf("agent id %v not found", id))
	}
	_, err := db.Exec("UPDATE jarvis.hosts SET online=TRUE WHERE systemId=?", id)
	return err
}

func (m *JarvisMysqlBackend) GrimReaper() {
	db := m.db
	// determine how fast the user could be notified after host offline
	checkInterval := 10 * time.Second
	for ;; time.Sleep(checkInterval) {
		result, err := db.Exec("UPDATE jarvis.hosts SET online=FALSE WHERE online=TRUE AND lastSeenAt < ? - INTERVAL 1 MINUTE", time.Now())
		if err != nil {
			log.Error(err.Error())
		}
		killed, err := result.RowsAffected()
		if err != nil {
			log.Error(err.Error())
		}
		log.WithField("killed", killed).Info("Online check")
	}
}

func (m *JarvisMysqlBackend) PostRegHost(h model.Host, hostId string) (err error) {
	// start a database transaction
	db := m.db
	result, err := db.Exec("UPDATE jarvis.hosts SET datacenter=?, rack=?, slot=?, tags=?, owner=?, osExpected=?, cpuExpected=?, memExpected=?, diskExpected=?, networkExpected=?, registered=TRUE, createdAt=?, updatedAt=? WHERE systemId=? and registered=FALSE ",
		h.DataCenter,
		h.Rack,
		h.Slot,
		helper.SafeMarshalJsonArray(h.Tags),
		h.Owner,
		helper.SafeMarshalJsonObj(h.OsExpected),
		helper.SafeMarshalJsonObj(h.CpuExpected),
		helper.SafeMarshalJsonObj(h.MemExpected),
		helper.SafeMarshalJsonArray(h.DiskExpected),
		helper.SafeMarshalJsonObj(h.NetworkExpected),
		time.Now(),
		time.Now(),
		hostId,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	log.WithFields(log.Fields{
		"row_affected": rows,
		"hostId": hostId,
	}).Info("post reg host")
	return nil
}

// factory method
func GetBackend() (*JarvisMysqlBackend, error) {
	mysqlConnectionLock.Lock()
	defer mysqlConnectionLock.Unlock()
	for ;b == nil;time.Sleep(2*time.Second) {
		dsn := fmt.Sprintf("root:passw0rd@tcp(%v)/jarvis?parseTime=true", options.MysqlAddr)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.WithError(err).Error("Open DB failed")
			continue
		}
		err = db.Ping()
		if err != nil {
			log.WithError(err).Error("Ping DB failed")
			continue
		}
		log.WithFields(log.Fields{
			"dsn": dsn,
		}).Info("MySQL connected.")
		b = &JarvisMysqlBackend{
			host: "localhost",
			port: 2379,
			db:   db,
		}
		//b.prepareStatements()
	}
	return b, nil
}
