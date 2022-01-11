package netx

import (
	"log"
	"testing"
)

// spy 2020/1/21

func TestLocalIP(t *testing.T) {
	log.Print(GetLocalIP())
}
