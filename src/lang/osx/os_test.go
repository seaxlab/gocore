package osx

import (
	"fmt"
	"testing"
)

// spy 2020/1/21

func TestSeaHome(t *testing.T) {
	fmt.Println(SeaHome())
}

func TestGetHostname(t *testing.T) {
	fmt.Printf("hostname=%s", GetHostname())
}

func TestGetWorkingDir(t *testing.T) {
	fmt.Println(GetWorkingDir())
}
