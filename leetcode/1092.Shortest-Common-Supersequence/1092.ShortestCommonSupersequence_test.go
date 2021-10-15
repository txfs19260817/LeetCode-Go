package _092_Shortest_Common_Supersequence

import (
	"testing"
)

func Test_shortestCommonSupersequence(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `str1 = "bbbaaaba", str2 = "bbababbb"`,
			args: args{"bbbaaaba", "bbababbb"},
			want: "bbabaaababb",
		},
		{
			name: `str1 = "abac", str2 = "cab"`,
			args: args{"abac", "cab"},
			want: "cabac",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCommonSupersequence(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("shortestCommonSupersequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
