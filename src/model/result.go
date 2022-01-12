package model

import "fmt"

// spy 2018-06-18

/** API返回值统一模型 */

type Result struct {
	Success bool        `json:"success"`
	TraceId string      `json:"traceId"`
	Code    string      `json:"code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Field   string      `json:"field,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccess() *Result {
	return &Result{
		Success: true,
		Data:    nil,
	}
}

func NewSuccessResult(data interface{}) *Result {
	return &Result{
		Success: true,
		Data:    data,
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

// -------------- method -------------------

// SetMsgF set msg info
func (r *Result) SetMsgF(format string, args ...interface{}) {
	r.Msg = fmt.Sprintf(format, args...)
}
