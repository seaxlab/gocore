package reflectx

// spy 2020/1/21

import (
	"reflect"
)

// Type Returns a string of variable type
func Type(data interface{}) string {
	return reflect.TypeOf(data).String()
}

// IsInt Int
func IsInt(data interface{}) bool {
	return kind(data) == reflect.Int

	//switch data.(type) {
	//case int:
	//	return true
	//default:
	//	return false
	//}
}

// IsInt64 Int64
func IsInt64(data interface{}) bool {
	return kind(data) == reflect.Int64
	//switch data.(type) {
	//case int64:
	//	return true
	//default:
	//	return false
	//}
}

// IsFloat32 check is float32
func IsFloat32(data interface{}) bool {
	return kind(data) == reflect.Float32
}

// IsFloat64 check is Float64
func IsFloat64(data interface{}) bool {
	return kind(data) == reflect.Float64
	//return Type(data) == "float64"
}

// IsString String
func IsString(data interface{}) bool {
	return kind(data) == reflect.String
}

// IsTime check is time.Time
func IsTime(data interface{}) bool {
	return Type(data) == "time.Time"
}

// IsBool check is bool
func IsBool(data interface{}) bool {
	return kind(data) == reflect.Bool
}

// IsTrue check is True
func IsTrue(data interface{}) bool {
	var flag = false
	switch Type(data) {
	case "string":
		flag = data == "true"
		break
	case "bool":
		flag = data == true
		break
	}
	return flag
}

// IsFalse check is false
func IsFalse(data interface{}) bool {
	return !IsTrue(data)
}

// IsSlice check is slice
// []int =true
// [5]int=false
func IsSlice(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Slice
}

// IsArray check is array
// []int =false
// [5]int=true
func IsArray(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Array
}

// IsArrayOrSlice check is array or slice
func IsArrayOrSlice(data interface{}) bool {
	kind := reflect.TypeOf(data).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

// IsMap check is map
func IsMap(data interface{}) bool {
	return reflect.TypeOf(data).Kind() == reflect.Map
}

//--------- private function ------
func kind(data interface{}) reflect.Kind {
	return reflect.TypeOf(data).Kind()
}
