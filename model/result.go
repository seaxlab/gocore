package model

// spy 2018-06-18

/** 对外返回基类 */
type Result struct {
	Success    bool        `json:"success"`
	ErrorCode  string      `json:"errorCode,omitempty"`
	ErrorMsg   string      `json:"errorMsg,omitempty"`
	ErrorField string      `json:"errorField,omitempty"`
	Data       interface{} `json:"data,omitempty"`
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

func NewFailResult(errorCode, errorMsg string) *Result {
	return &Result{
		Success:   false,
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
		Data:      nil,
	}
}
