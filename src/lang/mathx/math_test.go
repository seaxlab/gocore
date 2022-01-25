package mathx

import (
	"fmt"
	"testing"
)

// spy 2022/1/25

func TestRound(t *testing.T) {
	fmt.Println(Round(3.1415926, 2))
	fmt.Println(Round(3.1415926, 3))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(3.1415926))
	fmt.Println(Floor(3.1415926))
}

func TestCeil(t *testing.T) {
	fmt.Println(Round(3.1415926, 2))
	fmt.Println(Round(3.1415926, 3))
}
