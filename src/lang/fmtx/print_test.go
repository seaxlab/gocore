package fmtx

import (
	"testing"
)

// spy 2022/1/12
func TestPrintlnF(t *testing.T) {
	PrintlnF("a=%d,key=%s", 1, "2")

	// %v通用类型
	PrintlnF("a=%v,key=%v", 1, "2")
}
