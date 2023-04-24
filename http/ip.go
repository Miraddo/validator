package http

import (
	"net"
)

type HttpIp interface {
	IsValidIP(ip string) bool
	IsValidIPv4(ip string) bool
	IsValidIPv6(ip string) bool
	IsValidTCPPort(port string) bool
	IsValidUDPPort(port string) bool
	IsValidDomain(domain string) bool
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsValidIPv4(ip string) bool {
	return net.ParseIP(ip).To4() != nil
}

func IsValidIPv6(ip string) bool {
	return net.ParseIP(ip).To16() != nil
}

func IsValidTCPPort(port string) bool {
	_, err := net.LookupPort("tcp", port)
	return err == nil
}

func IsValidUDPPort(port string) bool {
	_, err := net.LookupPort("udp", port)
	return err == nil
}

func IsValidDomain(domain string) bool {
	_, err := net.LookupHost(domain)
	return err == nil
}
