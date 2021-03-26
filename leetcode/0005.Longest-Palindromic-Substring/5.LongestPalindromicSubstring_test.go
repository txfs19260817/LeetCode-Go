package _005_Longest_Palindromic_Substring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `s = "babad"`,
			args: args{"babad"},
			want: "bab",
		},
		{
			name: `s = "cbbd"`,
			args: args{"cbbd"},
			want: "bb",
		},
		{
			name: `s = "a"`,
			args: args{"a"},
			want: "a",
		},
		{
			name: `s = "ac"`,
			args: args{"ac"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
