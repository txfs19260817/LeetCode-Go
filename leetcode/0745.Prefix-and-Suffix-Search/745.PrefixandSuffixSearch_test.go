package leetcode

import (
	"testing"
)

func TestWordFilter(t *testing.T) {
	wordFilter := Constructor([]string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"})
	tests := []struct {
		args []string
		want int
	}{
		{
			args: []string{"bccbacbcba", "a"},
			want: 9,
		},
		{
			args: []string{"ab", "abcaccbcaa"},
			want: 4,
		},
		{
			args: []string{"a", "aa"},
			want: 5,
		},
		{
			args: []string{"cabaaba", "abaaaa"},
			want: 0,
		},
		{
			args: []string{"cacc", "accbbcbab"},
			want: 8,
		},
		{
			args: []string{"ccbcab", "bac"},
			want: 1,
		},
		{
			args: []string{"bac", "cba"},
			want: 2,
		},
		{
			args: []string{"ac", "accabaccaa"},
			want: 5,
		},
		{
			args: []string{"bcbb", "aa"},
			want: 3,
		},
		{
			args: []string{"ccbca", "cbcababac"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args[0]+"-"+tt.args[1], func(t *testing.T) {
			if got := wordFilter.F(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
