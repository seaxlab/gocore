package reflectx

import (
	"strings"
	"testing"
)

// spy 2022/1/26

func TestGetFuncName(t *testing.T) {
	name := GetFuncName(TestGetFuncName)
	t.Log(name)
	if !strings.HasSuffix(name, ".TestGetFuncName") {
		t.Error("get func name error")
	}
}
