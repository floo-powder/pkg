package handler

type BaseResponse struct {
	// Code http status code
	Code int `json:"code"`
	// Msg responce message
	Msg string `json:"msg"`
	// Data response data
	Data interface{} `json:"data"`
}

type Payload struct {
	// Value one row value
	Value []interface{} `json:"value"`
}

const (
	BaseURL = "/"
)
