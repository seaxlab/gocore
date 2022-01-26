package jsonx

import (
	"bytes"
	"encoding/json"
)

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

// ToIndentString converts the golang value to indent JSON string, such as a struct, map, slice, array and so on
func ToIndentString(obj interface{}) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
