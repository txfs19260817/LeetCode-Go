package _811_Subdomain_Visit_Count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `["9001 discuss.leetcode.com"]`,
			args: args{[]string{"9001 discuss.leetcode.com"}},
			want: []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"},
		},
		{
			name: `["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]`,
			args: args{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}},
			want: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, subdomainVisits(tt.args.cpdomains), tt.want)
		})
	}
}
