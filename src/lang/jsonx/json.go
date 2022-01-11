package jsonx

import "encoding/json"

// spy 2020/1/21

// ToJsonString Returns a Json string
func ToJsonString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// DecodeJsonString decode string to interface{}
func DecodeJsonString(s string) (interface{}, error) {
	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	return i, err
}
