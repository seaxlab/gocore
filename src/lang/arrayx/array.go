package arrayx

import (
	"math/rand"
	"time"
)

// spy 2020/5/1

func ToStrArray(data interface{}) []string {
	return data.([]string)
}

func Shuffle(data []interface{}) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
}

func ShuffleNew(data []interface{}) []interface{} {
	// clone first
	result := make([]interface{}, 0, len(data))
	for _, v := range data {
		result = append(result, v)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

// ShuffleInt Int
func ShuffleInt(data []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
}

// ShuffleIntNew an array
func ShuffleIntNew(data []int) []int {
	// clone first
	result := make([]int, 0, len(data))
	for _, v := range data {
		result = append(result, v)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

type MapFunc func(interface{}) interface{}

// Map 经过MapFunc算子进行计算
// Invokes `MapFunc` for each element in the array. Creates a new array containing the values returned by `MapFunc`
// Example:
//	func MyMapFunc(v interface{}) interface{} {
//		return v.(int) * 3
//	}
//	var myArray = []interface{}{1, 2, 3, 4}
//	result := arrayx.Map(myArray, MyMapFunc)  // [3, 6, 9, 12]
func Map(arr []interface{}, mapFunc MapFunc) []interface{} {
	if arr == nil || mapFunc == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		result = append(result, mapFunc(v))
	}
	return result
}

// Compact 去除nil元素
func Compact(arr []interface{}) []interface{} {
	if arr == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if v != nil {
			result = append(result, v)
		}
	}
	return result
}

//
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

// Merge 合并两个数组
func Merge(a, b []interface{}) []interface{} {
	return append(a, b...)
}

func MergeInt(slice1, slice2 []int) (c []int) {
	c = append(slice1, slice2...)
	return
}

func MergeInt64(slice1, slice2 []int64) (c []int64) {
	c = append(slice1, slice2...)
	return
}

func MergeString(slice1, slice2 []string) (c []string) {
	c = append(slice1, slice2...)
	return
}

// Unique 去重
// Returns a new array with duplicates removed.
// Example:
//	var myArray = []interface{}{1, 2, 3, 3, 4}
//	result := Unique(myArray)  // [1, 2, 3, 4]
func Unique(arr []interface{}) []interface{} {
	if arr == nil || len(arr) <= 1 {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if len(result) == 0 {
			result = append(result, v)
			continue
		}

		duplicate := false
		for _, r := range result {
			if r == v {
				duplicate = true
				break
			}
		}

		if !duplicate {
			result = append(result, v)
		}
	}
	return result
}
