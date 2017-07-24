package backend

import "git.oschina.net/k2ops/jarvis/server/api/model"

type JarvisBackend interface {
	CreateHost(h model.Host) error
	SearchHost(query map[string]string) ([]model.Host, error)
	UpdateHost(query map[string]string, h model.Host) error
	DeleteHost(query map[string]string) error
}


