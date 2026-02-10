package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpFirewall_MainExample(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.1.100"},
		{"DENY", "192.168.1.0/24"},
		{"ALLOW", "192.168.0.0/16"},
		{"DENY", "0.0.0.0/0"},
	})
	assert.True(t, fw.AllowAccess("192.168.1.100"))
	assert.False(t, fw.AllowAccess("192.168.1.50"))
	assert.True(t, fw.AllowAccess("192.168.2.10"))
	assert.False(t, fw.AllowAccess("10.0.0.1"))
}

func TestIpFirewall_DenyAll(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"DENY", "0.0.0.0/0"},
	})
	assert.False(t, fw.AllowAccess("1.2.3.4"))
	assert.False(t, fw.AllowAccess("255.255.255.255"))
	assert.False(t, fw.AllowAccess("0.0.0.0"))
}

func TestIpFirewall_ExactMatch32(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"ALLOW", "10.0.0.5/32"},
		{"DENY", "0.0.0.0/0"},
	})
	assert.True(t, fw.AllowAccess("10.0.0.5"))
	assert.False(t, fw.AllowAccess("10.0.0.6"))
}

// --------------- AllowCIDR follow-up tests ---------------

func TestAllowCIDR_EntireRangeCoveredByAllow(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.0.0/16"},
		{"DENY", "0.0.0.0/0"},
	})
	// /24 is fully inside the /16 ALLOW rule.
	assert.True(t, fw.AllowCIDR("192.168.1.0/24"))
}

func TestAllowCIDR_EntireRangeCoveredByDeny(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"DENY", "192.168.1.0/24"},
		{"ALLOW", "0.0.0.0/0"},
	})
	assert.False(t, fw.AllowCIDR("192.168.1.0/24"))
}

func TestAllowCIDR_TwoHalvesBothAllowed(t *testing.T) {
	// Two narrower ALLOW rules together cover the query, before catch-all DENY.
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.1.0/25"},
		{"ALLOW", "192.168.1.128/25"},
		{"DENY", "0.0.0.0/0"},
	})
	assert.True(t, fw.AllowCIDR("192.168.1.0/24"))
}

func TestAllowCIDR_OneHalfMissing(t *testing.T) {
	// Only the first half is ALLOW; second half falls through to DENY.
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.1.0/25"},
		{"DENY", "0.0.0.0/0"},
	})
	assert.False(t, fw.AllowCIDR("192.168.1.0/24"))
}

func TestAllowCIDR_MainExample(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.1.100"},
		{"DENY", "192.168.1.0/24"},
		{"ALLOW", "192.168.0.0/16"},
		{"DENY", "0.0.0.0/0"},
	})
	// /24 has 192.168.1.50 denied → not all-allowed.
	assert.False(t, fw.AllowCIDR("192.168.1.0/24"))
	// /16 minus /24 deny range — but the /24 deny is inside, so not all-allowed.
	assert.False(t, fw.AllowCIDR("192.168.0.0/16"))
	// A different /24 fully covered by /16 ALLOW.
	assert.True(t, fw.AllowCIDR("192.168.2.0/24"))
	// Outside 192.168 → DENY.
	assert.False(t, fw.AllowCIDR("10.0.0.0/8"))
}

func TestAllowCIDR_SingleIP32(t *testing.T) {
	fw := NewIpFirewall([][]string{
		{"ALLOW", "192.168.1.100"},
		{"DENY", "0.0.0.0/0"},
	})
	assert.True(t, fw.AllowCIDR("192.168.1.100/32"))
	assert.False(t, fw.AllowCIDR("192.168.1.50/32"))
}
