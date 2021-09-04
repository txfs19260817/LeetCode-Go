package _516_Longest_Palindromic_Subsequence

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "cbbd"`,
			args: args{"cbbd"},
			want: 2,
		},
		{
			name: `s = "a"`,
			args: args{"a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
			if got := longestPalindromeSubseq2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq2() = %v, want %v", got, tt.want)
			}
		})
	}
}
