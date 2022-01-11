package idx

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

func TestGetYYYYMMDDHHMMSS(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetYYYYMMDDHHMMSS())
	}
}

func TestGetYYYYMMDDHHMM(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetYYYYMMDDHHMM())
	}
}

func TestGetYYYYMMDDHH(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetYYYYMMDDHH())
	}
}

func TestGetYYYYMMDD(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", GetYYYYMMDD())
	}
}
