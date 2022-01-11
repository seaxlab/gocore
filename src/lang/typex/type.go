package typex

// spy 2020/1/21

import (
	"reflect"
)

// Type Returns a string of variable type
func Type(arg interface{}) string {
	return reflect.TypeOf(arg).String()
}

// IsInt Int
func IsInt(arg interface{}) bool {
	switch arg.(type) {
	case int:
		return true
	default:
		return false
	}
}

// IsInt64 Int64
func IsInt64(arg interface{}) bool {
	switch arg.(type) {
	case int64:
		return true
	default:
		return false
	}
}

// IsFloat Float64
func IsFloat(arg interface{}) bool {
	return Type(arg) == "float64"
}

// IsString String
func IsString(arg interface{}) bool {
	return Type(arg) == "string"
}

// IsTime check is time.Time
func IsTime(arg interface{}) bool {
	return Type(arg) == "time.Time"
}

// IsBool check is bool
func IsBool(arg interface{}) bool {
	return Type(arg) == "bool"
}

// IsTrue check is True
func IsTrue(arg interface{}) bool {
	var flag = false
	switch Type(arg) {
	case "string":
		flag = arg == "true"
		break
	case "bool":
		flag = arg == true
		break
	}
	return flag
}

// IsFalse check is false
func IsFalse(arg interface{}) bool {
	return !IsTrue(arg)
}

// IsSlice check is slice
func IsSlice(arg interface{}) bool {
	return reflect.TypeOf(arg).Kind().String() == "slice"
}

// IsMap check is map
func IsMap(arg interface{}) bool {
	return reflect.TypeOf(arg).Kind().String() == "map"
}
