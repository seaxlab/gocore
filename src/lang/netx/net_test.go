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

func TestIsPublicIPStr(t *testing.T) {
	log.Println(IsPublicIPStr("192.168.0.1"))
	log.Println(IsPublicIPStr("172.168.0.1"))
}

func TestIsPrivateIPStr(t *testing.T) {
	log.Println(IsPrivateIPStr("192.168.0.1"))
	log.Println(IsPrivateIPStr("172.168.0.1"))
}
