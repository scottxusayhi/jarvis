package backend

import "git.oschina.net/k2ops/jarvis/server/api/model"

type JarvisBackend interface {
	CreateHost(h model.Host)
	SearchHost (query map[string]string) []model.Host
	UpdateHost(query map[string]string, h model.Host)
	DeleteHost(query map[string]string)
}


