package _003_Longest_Substring_Without_Repeating_Characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "abcabcbb"`,
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: `s = "bbbbb"`,
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: `s = "pwwkew"`,
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: `s = "abba"`,
			args: args{"abba"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
