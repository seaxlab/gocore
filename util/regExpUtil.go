package util

import "regexp"

// spy 2020/4/15
func GetMatchValue(content string, key string) string {
	reg := regexp.MustCompile("(" + key + ")=([^,]*)")
	one := reg.FindAllStringSubmatch(content, -1)

	//fmt.Println(one) //[[app=user-22 app user-22]]

	if one == nil || len(one) == 0 {
		return ""
	}

	if len(one[0]) == 3 {
		//fmt.Println(one[0][2])
		return one[0][2]
	}
	return ""
}
