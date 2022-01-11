package arrayx

import (
	"fmt"
	"testing"
)

// spy 2022/1/11

func MyMapFunc(v interface{}) interface{} {
	return v.(int) * 3
}

func TestMap(t *testing.T) {

	var myArray = []interface{}{1, 2, 3, 4}
	result := Map(myArray, MyMapFunc) // [3, 6, 9, 12]

	fmt.Println(result)
}
