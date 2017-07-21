package mysql

import "git.oschina.net/k2ops/jarvis/server/api/model"

type JarvisMysqlBackend struct {
	host string
	port uint16
}

func (m *JarvisMysqlBackend) CreateHost(h model.Host) {
	panic("implement me")
}

func (m *JarvisMysqlBackend) SearchHost(query map[string]string) []model.Host {
	panic("implement me")
}

func (m *JarvisMysqlBackend) UpdateHost(query map[string]string, h model.Host) {
	panic("implement me")
}

func (m *JarvisMysqlBackend) DeleteHost(query map[string]string) {
	panic("implement me")
}
