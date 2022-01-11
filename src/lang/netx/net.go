package netx

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

// IsPublicIP 判断ip是否是公网IP
func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// IsIPv4 判断是否是IPv4
func IsIPv4(ip net.IP) bool {
	return ip.To4() != nil
}

// GetLanIPs 获取局域网本地IP地址列表
func GetLanIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok &&
			!ipnet.IP.IsLoopback() &&
			!ipnet.IP.IsLinkLocalMulticast() &&
			!ipnet.IP.IsLinkLocalUnicast() &&
			IsIPv4(ipnet.IP) &&
			!IsPublicIP(ipnet.IP) {

			ips = append(ips, ipnet.IP.String())
		}
	}

	return
}
