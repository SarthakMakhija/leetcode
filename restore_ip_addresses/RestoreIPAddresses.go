package restore_ip_addresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var restoreIpAddressesInner func(index int, ipPartCount int, currentIpFragment string, ip string) []string
	restoreIpAddressesInner = func(index int, ipPartCount int, currentIpFragment string, ip string) []string {
		intPart, _ := strconv.Atoi(currentIpFragment)
		switch {
		case len(currentIpFragment) != 1 && strings.Index(currentIpFragment, "0") == 0:
			return []string{}
		case intPart > 255:
			return []string{}
		case ipPartCount == 4 && index == len(s):
			return []string{ip[0 : len(ip)-1]}
		case index > len(s):
			return []string{}
		case ipPartCount > 4:
			return []string{}
		}
		var ips []string
		if index+1 <= len(s) {
			ips = append(ips, restoreIpAddressesInner(index+1, ipPartCount+1, s[index:index+1], ip+s[index:index+1]+".")...)
		}
		if index+2 <= len(s) {
			ips = append(ips, restoreIpAddressesInner(index+2, ipPartCount+1, s[index:index+2], ip+s[index:index+2]+".")...)
		}
		if index+3 <= len(s) {
			ips = append(ips, restoreIpAddressesInner(index+3, ipPartCount+1, s[index:index+3], ip+s[index:index+3]+".")...)
		}
		return ips
	}
	return restoreIpAddressesInner(0, 0, "", "")
}
