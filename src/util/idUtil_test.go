package util

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetUUID())
	}
}

func TestShortUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetShortUUID())
	}
}
