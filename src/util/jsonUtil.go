package util

import "encoding/json"

// spy 2020/1/21

// ToJsonString Returns a Json string
func ToJsonString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// DecodeJsonString decode string to interface{}
func DecodeJsonString(s string) interface{} {
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		return nil
	}
	return i
}
