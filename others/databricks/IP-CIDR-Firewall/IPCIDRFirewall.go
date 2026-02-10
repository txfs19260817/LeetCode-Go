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

// AllowCIDR returns true if EVERY IP in the given CIDR range would be
// allowed by the firewall rules.
//
// Uses recursive bisection: scan rules in priority order; if a single rule
// covers the entire query range, return its action.  If a higher-priority
// narrower rule partially overlaps the query, split the query into two
// halves (prefix+1) and check each recursively.
func (fw *IpFirewall) AllowCIDR(cidr string) bool {
	network, _ := parseCIDR(cidr)
	return fw.allowRange(network, parsePrefixLen(cidr))
}

func (fw *IpFirewall) allowRange(network uint32, prefixLen int) bool {
	qMask := prefixToMask(prefixLen)
	needsSplit := false

	for _, r := range fw.rules {
		// Does this rule cover the ENTIRE query range?
		// (rule is at least as broad, and query falls within rule)
		if qMask&r.mask == r.mask && network&r.mask == r.network {
			if needsSplit {
				break // higher-priority narrower rule exists — must recurse
			}
			return r.allow
		}
		// Is this rule a proper subset of the query? (strictly narrower & contained)
		if qMask&r.mask != r.mask && r.network&qMask == network {
			needsSplit = true
		}
	}

	// Base case: single IP — scan rules directly.
	if prefixLen >= 32 {
		for _, r := range fw.rules {
			if network&r.mask == r.network {
				return r.allow
			}
		}
		return false
	}

	// Split into two /prefixLen+1 halves and require both to be allowed.
	bit := uint32(1) << (31 - prefixLen)
	return fw.allowRange(network, prefixLen+1) &&
		fw.allowRange(network|bit, prefixLen+1)
}

// parsePrefixLen extracts the prefix length from a CIDR string (default 32).
func parsePrefixLen(cidr string) int {
	if idx := strings.Index(cidr, "/"); idx != -1 {
		p, _ := strconv.Atoi(cidr[idx+1:])
		return p
	}
	return 32
}

// prefixToMask converts a prefix length to a subnet mask.
func prefixToMask(prefixLen int) uint32 {
	if prefixLen == 0 {
		return 0
	}
	return ^uint32(0) << (32 - prefixLen)
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
