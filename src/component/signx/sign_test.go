package signx

import (
	"fmt"
	"github.com/seaxlab/gocore/src/lang/timex"
	"github.com/stretchr/testify/assert"
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
	//params["timestamp"] = "2021-01-01 12:30:40"
	params["timestamp"] = timex.Format(time.Now(), timex.FORMAT_YYYY_MM_DD_HH_MM_SS)

	appSecret := "abcexvsh12"
	sign := Get(appSecret, params)
	fmt.Printf("sign=%s", sign)

	assert.EqualValues(t, "C3A2A6CF1A875BEA3F4EEA10C36784BE", sign)
}
