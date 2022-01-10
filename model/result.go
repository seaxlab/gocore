package model

// spy 2018-06-18

/** 对外返回基类 */

type Result struct {
	Success bool        `json:"success"`
	TraceId string      `json:"traceId"`
	Code    string      `json:"code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Field   string      `json:"field,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResult() *Result {
	return &Result{
		Success: true,
		Data:    nil,
	}
}

func NewFail() *Result {
	return &Result{
		Success: false,
		Data:    nil,
	}
}

func NewFailResult(code, msg string) *Result {
	return &Result{
		Success: false,
		Code:    code,
		Msg:     msg,
		Data:    nil,
	}
}
