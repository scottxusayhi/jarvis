package helper

import (
	"encoding/json"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type pageInfo struct {
	Size      int `json:"size"`
	TotalSize int `json:"totalSize"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
}

func (p *pageInfo) Offset() int {
	return p.Page * p.PerPage
}

func (p *pageInfo) Limit() int {
	return p.PerPage
}

func (p *pageInfo) SetResult(size int, totalSize int, totalPage int) {
	p.Size = size
	p.TotalSize = totalSize
	p.TotalPage = totalPage
}

func NewPageInfo(perPage int, page int) pageInfo {
	return pageInfo{
		PerPage: perPage,
		Page:    page,
	}
}

func DefaultPageInfo() pageInfo {
	return pageInfo{
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

func WrapHostListResponse(code int, message string, hosts []model.Host, p pageInfo) ([]byte, error) {
	result := jsonObj{
		"code":     code,
		"message":  message,
		"pageInfo": p,
		"list":     hosts,
	}
	return json.Marshal(result)
}
