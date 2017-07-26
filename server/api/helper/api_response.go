package helper

import (
	"net/http"
	"git.oschina.net/k2ops/jarvis/server/api/model"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type pageInfo struct {
	Size uint `json:"size"`
	TotalSize uint `json:"totalSize"`
	TotalPage uint `json:"totalPage"`
	Page uint `json:"page"`
	PerPage uint `json:"perPage"`
}

func (p *pageInfo) Offset() uint {
	return p.Page * p.PerPage
}

func (p *pageInfo) Limit() uint {
	return p.PerPage
}

func (p *pageInfo) Result (size uint, totalSize uint, totalPage uint)  {
	p.Size = size
	p.TotalSize = totalSize
	p.TotalPage = totalPage
}

func NewPageInfo(perPage uint, page uint) pageInfo {
	return pageInfo{
		PerPage: perPage,
		Page: page,
	}
}

func DefaultPageInfo() pageInfo {
	return pageInfo{
		Page: 1,
		PerPage: 20,
	}
}

func Write400Error (w http.ResponseWriter, message string) {
	writeError(w, http.StatusBadRequest, message)
}

func Write500Error (w http.ResponseWriter, message string) {
	writeError(w, http.StatusInternalServerError, message)
}

func writeError (w http.ResponseWriter, code int, message string) {
	response := model.ApiResBody{}
	response.Code = 1
	response.Message = message
	bytes, err := json.Marshal(model.ApiResBody{
		Code: 1,
		Message: message,
	})
	if err != nil {
		log.Error(err.Error())
	}
	w.WriteHeader(code)
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
	result := jsonObj {
		"code": code,
		"message": message,
		"pageInfo": p,
		"list": hosts,
	}
	return json.Marshal(result)
}