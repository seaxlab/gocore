package sortx

import (
	"fmt"
	"testing"
)

// spy 2022/1/21
func TestName(t *testing.T) {
	var names = []string{"b", "a", "e", "c", "d"}
	StringAsc(names)
	fmt.Println(names)
}
