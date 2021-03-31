package _049_Group_Anagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: `strs = ["eat","tea","tan","ate","nat","bat"]`,
			args: args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: `strs = ["bd","aacc"]`,
			args: args{[]string{"bd", "aacc"}},
			want: [][]string{
				{"bd"},
				{"aacc"},
			},
		},
		{
			name: `strs = ["bbd","aacc"]`,
			args: args{[]string{"bbd", "aacc"}},
			want: [][]string{
				{"bbd"},
				{"aacc"},
			},
		},
		{
			name: `strs = ["may","max"]`,
			args: args{[]string{"may", "max"}},
			want: [][]string{
				{"may"},
				{"max"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) < len(got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return len(tt.want[i]) < len(tt.want[j])
			})
			for i := range got {
				assert.ElementsMatch(t, tt.want[i], got[i])
			}
		})
	}
}
