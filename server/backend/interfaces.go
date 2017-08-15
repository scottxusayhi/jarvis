package backend

import (
	"errors"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"net/url"
	"strings"
	"git.oschina.net/k2ops/jarvis/server/api/helper"
	"strconv"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var nonDataColumns []string = []string{
	"page",
	"perPage",
	"type",
}

type Query map[string]string

func (q Query) SqlString() string {
	var s []string
	index := 0
	for k, v := range q {
		if !contains(nonDataColumns, k) {
			s = append(s, k+"=\""+v+"\"")
			index += 1
		}
	}
	if len(s) > 0 {
		return "where " + strings.Join(s, " and ")
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
		// the first value appears takes effect if a query param has multiple value
		result[key] = value[0]
	}
	return result
}

func PageInfoFromUrlQuery(query url.Values) (info helper.PageInfo) {
	info = helper.DefaultPageInfo()
	for key, value := range query {
		if strings.EqualFold(key, "perPage") {
			perPage, err := strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				log.WithError(err).Error(fmt.Sprintf("invalid query %v: should be uint but got %v", key, value))
			}
			info.PerPage = int(perPage)
		}

		if strings.EqualFold(key, "page") {
			page, err := strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				log.WithError(err).Error(fmt.Sprintf("invalid query %v: should be uint but got %v", key, value))

			}
			info.Page = int(page)
		}
	}
	return
}

func PageInfo(query Query) (info helper.PageInfo) {
	info = helper.DefaultPageInfo()
	for key, value := range query {
		if strings.EqualFold(key, "perPage") {
			perPage, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.WithError(err).Error(fmt.Sprintf("invalid query %v: should be uint but got %v", key, value))
			}
			info.PerPage = int(perPage)
		}

		if strings.EqualFold(key, "page") {
			page, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.WithError(err).Error(fmt.Sprintf("invalid query %v: should be uint but got %v", key, value))

			}
			info.Page = int(page)
		}
	}
	return
}


type JarvisBackend interface {
	CreateHost(h model.Host) error
	SearchHost(q Query) ([]model.Host, error)
	UpdateHost(q Query, h model.Host) error
	DeleteHost(q Query) error
}

var ErrHostExist = errors.New("host already exists")
var ErrHostNotFound = errors.New("host not found")
