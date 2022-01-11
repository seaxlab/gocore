package netx

import (
	"log"
	"net"
	"testing"
)

// spy 2020/1/21

func TestLocalIP(t *testing.T) {
	log.Print(GetLocalIP())
}

func TestIsPublicIP(t *testing.T) {
	log.Println(IsPublicIP(net.IPv4(127, 0, 0, 1)))
	log.Println(IsPublicIP(net.IPv4(115, 206, 113, 54)))
}
