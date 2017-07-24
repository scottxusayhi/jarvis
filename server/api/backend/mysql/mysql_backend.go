package mysql

import (
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	backend *JarvisMysqlBackend = nil
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
	
	m.stmtInsertHost, err = db.Prepare("insert into hosts values(?);")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (m *JarvisMysqlBackend) CreateHost(h model.Host) error {
	log.Info("create host")

	panic("implement me")
}

func (m *JarvisMysqlBackend) SearchHost(query map[string]string) ([]model.Host, error) {
	hosts := make([]model.Host, 10)
	db := m.db
	result, err := db.Exec("select * from hosts;")
	if err != nil {
		log.Error(err.Error())
	}
	log.Info(result)
	return hosts, nil
}

func (m *JarvisMysqlBackend) GetOneHost(dc string, rack string, slot string, hostname string) (*model.Host, error) {
	var err error
	host := &model.Host{}
	err = m.stmtGetOneHost.QueryRow(dc, rack, slot, hostname).Scan(
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
		return nil, err
	}
	return host, nil
}

func (m *JarvisMysqlBackend) UpdateHost(query map[string]string, h model.Host) error {
	panic("implement me")
}

func (m *JarvisMysqlBackend) DeleteHost(query map[string]string) error {
	panic("implement me")
}


func GetBackend () (*JarvisMysqlBackend, error) {
	if backend == nil {
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
		backend = &JarvisMysqlBackend{
			host: "localhost",
			port: 2379,
			db: db,
		}
		backend.prepareStatements()
	}
	return backend, nil
}