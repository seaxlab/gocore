package model

import (
	"encoding/json"
	"fmt"
	"github.com/seaxlab/gocore/src/lang/jsonx"
	"testing"
)

// spy 2022/1/10
func TestFailResult(t *testing.T) {
	result := NewFailResult("1", "a")
	fmt.Println(json.Marshal(result))
	fmt.Println(jsonx.ToJsonString(result))
	result.SetMsgF("abc%v%v", 1, 2)

	str, _ := jsonx.ToJsonString(result)
	fmt.Println(str)
}
