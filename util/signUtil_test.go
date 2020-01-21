package util

import (
	"testing"
	"time"
)

// spy 2020/1/21

func TestSignUtil(t *testing.T) {

	params := map[string]interface{}{}
	params["method"] = "get"
	params["app_key"] = "abc123456"
	params["format"] = "json"
	params["v"] = "1.0"
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")

	appSecret := "abcexvsh12"
	logger.Info("sign is", GetSign(appSecret, params))
}
