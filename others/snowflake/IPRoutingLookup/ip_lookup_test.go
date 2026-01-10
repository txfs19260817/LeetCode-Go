package snowflake

import (
	"reflect"
	"testing"
)

func Test_routeLookup(t *testing.T) {
	type args struct {
		routes  [][]string
		queries []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1: Basic Overlap",
			args: args{
				routes: [][]string{
					{"10.0.0.0/8", "a.com"},
					{"10.1.0.0/16", "b.com"},
					{"10.1.2.0/24", "c.com"},
				},
				queries: []string{"10.1.2.7", "10.2.3.4", "192.168.1.1"},
			},
			want: []string{"c.com", "a.com", ""},
		},
		{
			name: "Case 2: Private Networks",
			args: args{
				routes: [][]string{
					{"192.168.1.0/24", "url1.com"},
					{"192.168.0.0/16", "url2.com"},
				},
				queries: []string{"192.168.1.123", "192.168.2.10"},
			},
			want: []string{"url1.com", "url2.com"},
		},
		{
			name: "Case 3: Nested with Default",
			args: args{
				routes: [][]string{
					{"0.0.0.0/0", "default.com"},
					{"10.0.0.0/8", "corporate.com"},
					{"10.1.1.0/24", "subnet.com"},
					{"10.1.1.100/30", "server_farm.com"},
					{"10.1.1.101/32", "critical_server.com"},
				},
				queries: []string{
					"10.1.1.101",
					"10.1.1.100",
					"10.1.1.104",
					"10.2.0.1",
					"8.8.8.8",
				},
			},
			want: []string{
				"critical_server.com",
				"server_farm.com",
				"subnet.com",
				"corporate.com",
				"default.com",
			},
		},
		{
			name: "Case 4: Complex Subnetting",
			args: args{
				routes: [][]string{
					{"192.168.1.0/24", "net_main.com"},
					{"192.168.1.0/25", "net_first_half.com"},
					{"192.168.1.128/25", "net_second_half.com"},
					{"192.168.1.240/28", "net_sixteenth.com"},
					{"192.168.1.254/31", "redundant_pair.com"},
					{"192.168.0.0/23", "supernet.com"},
				},
				queries: []string{
					"192.168.1.0",
					"192.168.1.127",
					"192.168.1.128",
					"192.168.1.240",
					"192.168.1.254",
					"192.168.1.255",
					"192.168.0.1",
				},
			},
			want: []string{
				"net_first_half.com",
				"net_first_half.com",
				"net_second_half.com",
				"net_sixteenth.com",
				"redundant_pair.com",
				"redundant_pair.com",
				"supernet.com",
			},
		},
		{
			name: "Case 5: Broad Ranges",
			args: args{
				routes: [][]string{
					{"0.0.0.0/1", "zero_half.com"},
					{"128.0.0.0/1", "one_half.com"},
					{"64.0.0.0/2", "quarter_two.com"},
					{"128.0.0.0/9", "bit_pattern.com"},
					{"1.1.1.0/24", "dns_net.com"},
					{"169.254.0.0/16", "link_local.com"},
				},
				queries: []string{
					"0.0.0.1",
					"64.0.0.1",
					"127.255.255.255",
					"128.0.0.1",
					"128.1.0.1",
					"1.1.1.1",
					"169.254.1.1",
					"192.0.2.1",
				},
			},
			want: []string{
				"zero_half.com",
				"quarter_two.com",
				"quarter_two.com",
				"bit_pattern.com",
				"bit_pattern.com",
				"dns_net.com",
				"link_local.com",
				"one_half.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := routeLookup(tt.args.routes, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("routeLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
