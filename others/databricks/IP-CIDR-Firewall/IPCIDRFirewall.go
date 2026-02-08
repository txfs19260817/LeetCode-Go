package databricks

import (
	"strconv"
	"strings"
)

type rule struct {
	allow   bool
	network uint32
	mask    uint32
}

// IpFirewall stores an ordered list of CIDR-based allow/deny rules.
type IpFirewall struct {
	rules []rule
}

// NewIpFirewall creates a firewall from an ordered slice of [action, cidr] pairs.
func NewIpFirewall(rules [][]string) *IpFirewall {
	fw := &IpFirewall{}
	for _, r := range rules {
		action, cidr := r[0], r[1]
		allow := action == "ALLOW"
		network, mask := parseCIDR(cidr)
		fw.rules = append(fw.rules, rule{allow: allow, network: network, mask: mask})
	}
	return fw
}

// AllowAccess returns true if ip is allowed by the first matching rule.
func (fw *IpFirewall) AllowAccess(ip string) bool {
	ipVal := parseIP(ip)
	for _, r := range fw.rules {
		if ipVal&r.mask == r.network {
			return r.allow
		}
	}
	// Should not reach here per problem guarantees.
	return false
}

// parseIP converts an IPv4 dotted-decimal string to a uint32.
func parseIP(ip string) uint32 {
	parts := strings.Split(ip, ".")
	var result uint32
	for _, p := range parts {
		v, _ := strconv.Atoi(p)
		result = (result << 8) | uint32(v)
	}
	return result
}

// parseCIDR converts a CIDR string (or plain IP) into a network address and mask.
func parseCIDR(cidr string) (network uint32, mask uint32) {
	prefixLen := 32
	ipStr := cidr
	if idx := strings.Index(cidr, "/"); idx != -1 {
		ipStr = cidr[:idx]
		prefixLen, _ = strconv.Atoi(cidr[idx+1:])
	}
	if prefixLen == 0 {
		mask = 0
	} else {
		mask = ^uint32(0) << (32 - prefixLen)
	}
	ip := parseIP(ipStr)
	network = ip & mask
	return
}
