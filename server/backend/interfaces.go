package backend

import (
	"errors"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"net/url"
	"strings"
)

var nonDbColumns []string = []string{
	"page",
	"perPage",
	"type",
}

type Query map[string]string

func (q Query) SqlString() string {
	var s []string
	index := 0
	for k, v := range q {
		if !contains(nonDbColumns, k) {
			s = append(s, k+"=\""+v+"\"")
			index += 1
		}
	}
	if len(s) > 0 {
		return " where " + strings.Join(s, " and ")
	}
	return ""
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}

func FromURLQuery(query url.Values) Query {
	result := Query{}
	for key, value := range query {
		// use the first value
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
