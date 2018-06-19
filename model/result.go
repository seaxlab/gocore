package model

// spy 2018-06-18

/** 对外返回基类 */
type Result struct {
	Success    bool        "json:success"
	ErrorCode  string      "json:errorCode"
	ErrorMsg   string      "json:errorMsg"
	ErrorField string      "json:errorField"
	Data       interface{} "json:data"
}

func NewResult() *Result {
	return &Result{
		Success: true,
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
