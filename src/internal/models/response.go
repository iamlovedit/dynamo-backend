package models

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HttpMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
