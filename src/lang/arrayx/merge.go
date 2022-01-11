package arrayx

// spy 2022/1/11

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
