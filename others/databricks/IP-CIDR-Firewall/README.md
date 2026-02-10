# IP CIDR Firewall

Design an IP firewall that determines whether an IPv4 address is allowed or denied based on an ordered list of "ALLOW" or "DENY" rules. Each rule specifies either a single IP or a CIDR block, and only the first matching rule decides the result.

**IpFirewall(rules [][]string)** - Initializes a firewall with `[action, cidr]` rules in priority order.
- `action` is `"ALLOW"` or `"DENY"`
- `cidr` is either an IP (e.g., `"1.2.3.4"`) or a CIDR (e.g., `"192.168.0.0/16"`)

**AllowAccess(ip string) bool** - Returns `true` if `ip` is allowed by the first matching rule; `false` otherwise. Guaranteed to match at least one rule.

## Example

**Input:**
```
["IpFirewall", "allowAccess", "allowAccess", "allowAccess", "allowAccess"]
[[["ALLOW", "192.168.1.100"], ["DENY", "192.168.1.0/24"], ["ALLOW", "192.168.0.0/16"], ["DENY", "0.0.0.0/0"]], ["192.168.1.100"], ["192.168.1.50"], ["192.168.2.10"], ["10.0.0.1"]]
```

**Output:**
```
[null, true, false, true, false]
```

**Explanation:**
- `"192.168.1.100"` → ALLOW (exact match rule 1)
- `"192.168.1.50"` → DENY (matches /24 rule 2)
- `"192.168.2.10"` → ALLOW (matches /16 rule 3)
- `"10.0.0.1"` → DENY (matches 0.0.0.0/0 rule 4)

## Follow-up: AllowCIDR

**AllowCIDR(cidr string) bool** — Return `true` if **every** IP in the given CIDR range would be allowed.

Uses recursive bisection: scan rules in priority order; if one rule covers the entire query, return its action. If a higher-priority narrower rule partially overlaps the query, split the query into two halves (prefix+1) and check each recursively.

```
fw.AllowCIDR("192.168.1.0/24") → false  // 192.168.1.50 is denied
fw.AllowCIDR("192.168.2.0/24") → true   // fully inside /16 ALLOW
```
