package arrayx

import (
	"fmt"
	"testing"
)

// spy 2022/1/11

func TestMerge(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	z := append(x, y...)
	fmt.Println(z)
}

// 还不如上面的简洁
func TestMerge2(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	// []T to an []interface{}
	// https://go.dev/doc/faq#convert_slice_of_interface
	x2 := make([]interface{}, len(x))
	for i, v := range x {
		x2[i] = v
	}

	y2 := make([]interface{}, len(y))
	for i, v := range y {
		y2[i] = v
	}

	a := Merge(x2, y2)
	fmt.Println(a)
}

func TestMergeString(t *testing.T) {
	x := []string{"a", "b"}
	y := []string{"c"}

	result := MergeString(x, y)
	fmt.Println(result)
}
