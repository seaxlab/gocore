package jsonx

import (
	"encoding/json"
	"fmt"
	"github.com/seaxlab/gocore/src/model"
	"testing"
)

// spy 2022/1/11
func TestJSON(t *testing.T) {
	result := model.NewFailResult("1", "a")
	fmt.Println(json.Marshal(result))
	fmt.Println(ToJsonString(result))
}
