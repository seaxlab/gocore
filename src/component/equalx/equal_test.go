package equalx

import (
	"log"
	"testing"
)

// spy 2022/1/11

func TestEqual(t *testing.T) {
	log.Println(IsEqInt16(1, 2))
	log.Println(IsEqStr("1", "1"))
}
