package model


// common keys that are attached to each api response body
type ApiResBody struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
