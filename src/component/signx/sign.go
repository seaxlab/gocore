package signx

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

// spy 2020/1/21

func Get(clientSecret string, params map[string]interface{}) string {
	var keys []string
	for k := range params {
		if k != "sign" && k != "access_token" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys) // asc

	signStr := clientSecret
	for _, key := range keys {
		signStr += key + getString(params[key])
	}
	signStr += clientSecret
	return md5Hash(signStr)
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}

func md5Hash(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToLower(hex.EncodeToString(cipherStr))
}
