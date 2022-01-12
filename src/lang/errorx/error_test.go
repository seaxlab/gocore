package errorx

import (
	"fmt"
	"testing"
)

// spy 2022/1/12
func TestNewF(t *testing.T) {
	err := NewF("error is %v, %v", 1, "plz contact admin")
	fmt.Println(err)
}
