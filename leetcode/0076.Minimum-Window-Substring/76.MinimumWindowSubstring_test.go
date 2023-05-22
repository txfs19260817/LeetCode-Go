package leetcode

import (
	"fmt"
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: fmt.Sprintf("s: %s, t: %s", "ADOBECODEBANC", "ABC"),
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "a", "a"),
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "a", "b"),
			args: args{
				s: "a",
				t: "b",
			},
			want: "",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "ab", "A"),
			args: args{
				s: "ab",
				t: "A",
			},
			want: "",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "ABC", "AC"),
			args: args{
				s: "ABC",
				t: "AC",
			},
			want: "ABC",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "bbaac", "aba"),
			args: args{
				s: "bbaac",
				t: "aba",
			},
			want: "baa",
		},
		{
			name: fmt.Sprintf("s: %s, t: %s", "babb", "baba"),
			args: args{
				s: "babb",
				t: "baba",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
