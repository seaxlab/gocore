package arrayx

// spy 2022/1/11

//TODO FIXME
func merge(a, b []interface{}) []interface{} {
	return append([]interface{}{}, append(a, b...)...)
}
