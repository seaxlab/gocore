package arrayx

import (
	"fmt"
	"testing"
)

// spy 2022/1/11

func TestMerge(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	z := append([]int{}, append(x, y...)...)
	//z := append([]int{}, append(x, y...)...)
	fmt.Println(z)
}
