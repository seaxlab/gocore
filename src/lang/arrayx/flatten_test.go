package arrayx

import (
	"log"
	"testing"
)

func TestFlatten(t *testing.T) {
	arr1 := []interface{}{1, 2, 3, 4}    // [1, 2, 3, 4]
	arr2 := []interface{}{5, 6, 7, arr1} // [5, 6, 7, [1, 2, 3, 4]]
	result := Flatten(arr2)              // [5, 6, 7, 1, 2, 3, 4]

	log.Println(result)
}
