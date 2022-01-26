package netx

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

// spy 2020/1/21
var (
	nativeEndian binary.ByteOrder

	ipv4PrivateCIDRString = []string{
		"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16",
		"172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.88.99.0/24", "192.168.0.0/16",
		"198.18.0.0/15", "198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/4", "240.0.0.0/4", "255.255.255.255/32",
	}
	ipv6PrivateCIDRString = []string{
		"::1/128", "::/128", "64:ff9b::/96", "::ffff:0:0/96", "100::/64", "2001::/23",
		"2001::/32", "2001:2::/48", "2001:db8::/32", "2001:10::/28", "2002::/16", "fc00::/7", "fe80::/10",
		"2001:20::/28", "ff00::/8",
	}

	ipv4PrivateCIDR []*net.IPNet
	ipv6PrivateCIDR []*net.IPNet
)

func init() {
	if nativeEndian == nil {
		var x uint32 = 0x01020304
		if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
			nativeEndian = binary.BigEndian
		} else {
			nativeEndian = binary.LittleEndian
		}
	}

	for _, v := range ipv4PrivateCIDRString {
		_, n, _ := net.ParseCIDR(v)
		ipv4PrivateCIDR = append(ipv4PrivateCIDR, n)
	}

	for _, v := range ipv6PrivateCIDRString {
		_, n, _ := net.ParseCIDR(v)
		ipv6PrivateCIDR = append(ipv6PrivateCIDR, n)
	}
}

// GetRemoteIP get remote ip.
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

// IsPublicIPStr 判断是否是公网IP
func IsPublicIPStr(ip string) bool {
	addr := net.ParseIP(ip)
	if addr == nil {
		return false
	}
	return IsPublicIP(addr)
}

// IsPublicIP 判断ip是否是公网IP
func IsPublicIP(addr net.IP) bool {
	if addr.IsLoopback() || addr.IsLinkLocalMulticast() || addr.IsLinkLocalUnicast() {
		return false
	}

	if addr.To4() != nil {
		for _, v := range ipv4PrivateCIDR {
			if v.Contains(addr) {
				return false
			}
		}
	} else {
		for _, v := range ipv6PrivateCIDR {
			if v.Contains(addr) {
				return false
			}
		}
	}
	// default true
	return true
}

func IsPrivateIPStr(ip string) bool {
	addr := net.ParseIP(ip)
	if addr == nil {
		return false
	}
	return IsPrivateIP(addr)
}

// IsPrivateIP 判断ip是否是内网地址
func IsPrivateIP(addr net.IP) bool {
	if addr.IsLoopback() || addr.IsMulticast() || addr.IsLinkLocalMulticast() || addr.IsLinkLocalUnicast() {
		return true
	}

	if addr.To4() != nil {
		for _, v := range ipv4PrivateCIDR {
			if v.Contains(addr) {
				return true
			}
		}
	} else {
		for _, v := range ipv6PrivateCIDR {
			if v.Contains(addr) {
				return true
			}
		}
	}

	return false // default false
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

// IPv4ToU32 convert ipv4(a.b.c.d) to uint32 in host byte order
func IPv4ToU32(ip net.IP) uint32 {
	if ip == nil {
		return 0
	}
	a := uint32(ip[12])
	b := uint32(ip[13])
	c := uint32(ip[14])
	d := uint32(ip[15])
	return uint32(a<<24 | b<<16 | c<<8 | d)
}

// U32ToIPv4 convert uint32 to ipv4(a.b.c.d) in host byte order
func U32ToIPv4(ip uint32) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}

// IPv4StrToU32 convert IPv4 string to uint32 in host byte order
func IPv4StrToU32(s string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(s)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

// U32ToIPv4Str convert uint32 to IPv4 string in host byte order
func U32ToIPv4Str(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
