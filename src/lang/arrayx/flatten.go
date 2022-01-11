package arrayx

// spy 2022/1/11

func Flatten(arr []interface{}) []interface{} {
	if arr == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		switch v.(type) {
		case []interface{}:
			fl := Flatten(v.([]interface{}))
			for _, f := range fl {
				result = append(result, f)
			}
		default:
			result = append(result, v)
		}
	}
	return result
}
