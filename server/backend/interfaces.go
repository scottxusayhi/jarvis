package backend

import (
	"errors"
	"github.com/scottxusayhi/jarvis/server/api/model"
	"net/url"
	"strings"
	"github.com/scottxusayhi/jarvis/server/api/helper"
	"strconv"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var nonDataColumns []string = []string{
	"page",
	"perPage",
	"type",
	"order",
}

type Query map[string]string

func (q Query) SqlStringWhere() string {
	var s []string
	index := 0
	for k, v := range q {
		if !contains(nonDataColumns, k) {
			var part string
			switch k {
			// tags query is special case, we should translate tags=tag1,tag2 to (JSON_CONTAINS(tags, JSON_ARRAY(tag1)) or JSON_CONTAINS(tags, JSON_ARRAY("tag2")))
			case "tags":
				values := strings.Split(v, ",")
				tempValues := []string{}
				for _, value := range values {
					tempValues = append(tempValues, fmt.Sprintf("JSON_CONTAINS(tags, JSON_ARRAY(\"%v\"))", value))
				}
				part = fmt.Sprintf("(%v)", strings.Join(tempValues, " or "))
				break
			default:
				// v: comma separated strings, e.g., k2data,k2data-plus
				// and we should parse it to ("k2data", "k2data-plus") for "where in" clause
				values := strings.Split(v, ",")
				tempValues := []string{}
				for _, value := range values {
					tempValues = append(tempValues, fmt.Sprintf("\"%v\"", value))
				}
				part = fmt.Sprintf("%v IN (%v)", k, strings.Join(tempValues, ","))
			}
			s = append(s, part)
			index += 1
		}
	}
	if len(s) > 0 {
		return " WHERE " + strings.Join(s, " AND ")
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

func SqlStringOrder(query Query) (string) {
	param := query["order"]
	if len(param)==0 {
		return ""
	}

	// orderby
	value := param
	var order string
	if strings.HasPrefix(param, "+") {
		order = "ASC"
		value = strings.TrimPrefix(param, "+")
	}
	if strings.HasPrefix(param, "-") {
		order = "DESC"
		value = strings.TrimPrefix(param, "-")
	}
	return fmt.Sprintf(" ORDER BY %v %v", value, order)
}


type JarvisBackend interface {
	CreateHost(h model.Host) error
	SearchHost(q Query) ([]model.Host, error)
	UpdateHost(q Query, h model.Host) error
	DeleteHost(q Query) error
}

var ErrHostExist = errors.New("host already exists")
var ErrHostNotFound = errors.New("host not found")
