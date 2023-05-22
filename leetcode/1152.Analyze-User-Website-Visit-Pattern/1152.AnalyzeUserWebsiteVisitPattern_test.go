package leetcode

import (
	"reflect"
	"testing"
)

func Test_mostVisitedPattern(t *testing.T) {
	type args struct {
		username  []string
		timestamp []int
		website   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "zkiikgv",
			args: args{[]string{"zkiikgv", "zkiikgv", "zkiikgv", "zkiikgv"}, []int{436363475, 710406388, 386655081, 797150921}, []string{"wnaaxbfhxp", "mryxsjc", "oz", "wlarkzzqht"}},
			want: []string{"oz", "mryxsjc", "wlarkzzqht"},
		},
		{
			name: "h",
			args: args{[]string{"h", "eiy", "cq", "h", "cq", "txldsscx", "cq", "txldsscx", "h", "cq", "cq"}, []int{527896567, 334462937, 517687281, 134127993, 859112386, 159548699, 51100299, 444082139, 926837079, 317455832, 411747930}, []string{"hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "hibympufi", "yljmntrclw", "hibympufi", "yljmntrclw"}},
			want: []string{"hibympufi", "hibympufi", "yljmntrclw"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostVisitedPattern(tt.args.username, tt.args.timestamp, tt.args.website); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisitedPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
