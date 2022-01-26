package reflectx

import (
	"reflect"
	"runtime"
)

// spy 2022/1/26

// GetFuncName get function name
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
