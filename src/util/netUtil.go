package util

import (
	"net"
	"net/http"
	"strings"
)

// spy 2020/1/21

// get remote ip.
func GetRemoteIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")

	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}

	if ip == "" {
		ip = strings.SplitN(r.RemoteAddr, ":", 2)[0]

		// Check localhost
		if ip == "[" {
			ip = "127.0.0.1"
		}
	}

	return ip
}

// GetLocalIP returns the non loop back local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
