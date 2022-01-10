package util

import (
	"strconv"
)

// spy 2020/1/21

func GetBoolean(mapData map[string]string, key string, defaultValue bool) bool {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseBool(data)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

func GetFloat64(mapData map[string]string, key string, defaultValue float64) float64 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseFloat(data, 64)
		if err == nil {
			return parsed
		}
	}

	return defaultValue
}

func GetFloat32(mapData map[string]string, key string, defaultValue float32) float32 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseFloat(data, 32)
		if err == nil {
			return float32(parsed)
		}
	}

	return defaultValue
}

func GetString(mapData map[string]string, key string, defaultValue string) string {
	data, present := mapData[key]
	if present {
		return data
	}

	return defaultValue
}

func GetInt64(mapData map[string]string, key string, defaultValue int64) int64 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseInt(data, 10, 64)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

func GetInt(mapData map[string]string, key string, defaultValue int) int {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseInt(data, 10, 32)
		if err == nil {
			return int(parsed)
		}
	}
	return defaultValue
}

func Contains(mapData map[string]string, key string) bool {
	_, present := mapData[key]
	return present
}

//val := reflect.Indirect(reflect.ValueOf(v))
//func DeleteMapItem(srcMap interface{}, args ...string) {
//	checkIsMap(srcMap)
//
//	ma := srcMap.(map[string]interface{})
//	for _, v := range args {
//		delete(ma, v)
//	}
//}

//func checkIsMap(srcMap interface{}) {
//	if IsMap(srcMap) {
//
//	} else {
//		panic("only support for map.")
//	}
//}
