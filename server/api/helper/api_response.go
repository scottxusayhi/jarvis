package helper

import (
	"encoding/json"
	"github.com/scottxusayhi/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"fmt"
)

type PageInfo struct {
	Size      int `json:"size"`
	TotalSize int `json:"totalSize"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
}

func (p *PageInfo) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *PageInfo) Limit() int {
	return p.PerPage
}

func (p *PageInfo) SqlString() string {
	return fmt.Sprintf(" LIMIT %v, %v", p.Offset(), p.Limit())
}

func (p *PageInfo) SetResult(size int, totalSize int, totalPage int) {
	p.Size = size
	p.TotalSize = totalSize
	p.TotalPage = totalPage
}

func (p *PageInfo) CalcTotalPage() {
	if p.TotalSize%p.PerPage > 0 {
		p.TotalPage = p.TotalSize/p.PerPage + 1
	} else {
		p.TotalPage = p.TotalSize / p.PerPage
	}
}

func NewPageInfo(perPage int, page int) PageInfo {
	return PageInfo{
		PerPage: perPage,
		Page:    page,
	}
}

func DefaultPageInfo() PageInfo {
	return PageInfo{
		Page:    1,
		PerPage: 20,
	}
}

func Write400Error(w http.ResponseWriter, message string) {
	WriteResponse(w, http.StatusBadRequest, 1, message)
}

func Write500Error(w http.ResponseWriter, message string) {
	WriteResponse(w, http.StatusInternalServerError, 1, message)
}

func WriteResponse(w http.ResponseWriter, httpCode int, apiCode int, message string) {
	bytes, err := json.Marshal(model.ApiResBody{
		Code:    apiCode,
		Message: message,
	})
	if err != nil {
		log.Error(err.Error())
	}
	w.WriteHeader(httpCode)
	w.Write(bytes)
}

func WrapResponse(src []byte, code int, message string) ([]byte, error) {
	jsonMap := map[string]interface{}{}
	err := json.Unmarshal(src, &jsonMap)
	if err != nil {
		return nil, err
	}
	jsonMap["code"] = code
	jsonMap["message"] = message
	return json.Marshal(jsonMap)
}

func WrapResponseSuccess(src []byte) ([]byte, error) {
	return WrapResponse(src, 0, "")
}

type jsonObj map[string]interface{}

func WrapHostListResponse(code int, message string, hosts []model.Host, p PageInfo) ([]byte, error) {
	result := jsonObj{
		"code":     code,
		"message":  message,
		"pageInfo": p,
		"list":     hosts,
	}
	return json.Marshal(result)
}
