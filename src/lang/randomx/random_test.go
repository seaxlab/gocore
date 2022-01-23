package randomx

import (
	"fmt"
	"testing"
)

// spy 2022/1/21
func TestString(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(String(8))
	}
}

func TestAlphabet(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(Alphabet(8))
	}
}

func TestNumber(t *testing.T) {
	for i := 0; i < 20; i++ {
		nums := Number(10, 1000, 10)
		fmt.Println(nums)
	}
}
