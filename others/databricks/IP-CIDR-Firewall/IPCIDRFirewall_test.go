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
