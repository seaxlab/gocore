package fmtx

import "fmt"

// spy 2022/1/12

func PrintlnF(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
