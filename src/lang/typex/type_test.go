package typex

import (
	"log"
	"testing"
)

// spy 2020/1/21

func TestIsBool(t *testing.T) {
	log.Println(IsBool("true"))
	log.Println(IsTrue("true"))
	log.Println(IsBool("true"))
}
