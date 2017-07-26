package backend

import (
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"strings"
	"errors"
	"net/url"
)

type Query map[string]string

func (q Query) SqlString() string  {
	count := len(q)
	s := make([]string, count)
	index := 0
	for k, v := range q {
		s[index] = k+"=\""+v+"\""
		index+=1
	}
	return strings.Join(s, " and ")
}

func FromURLQuery (query url.Values) Query {
	result := Query{}
	for key, value := range query {
		result[key] = value[0]
	}
	return result
}

type JarvisBackend interface {
	CreateHost(h model.Host) error
	SearchHost(q Query) ([]model.Host, error)
	UpdateHost(q Query, h model.Host) error
	DeleteHost(q Query) error
}

var ErrHostExist = errors.New("host already exists")
var ErrHostNotFound = errors.New("host not found")


