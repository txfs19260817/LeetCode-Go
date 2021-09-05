package _143_Longest_Common_Subsequence

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `text1 = "abcde", text2 = "ace"`,
			args: args{"abcde", "ace"},
			want: 3,
		},
		{
			name: `text1 = "ace", text2 = "ace"`,
			args: args{"ace", "ace"},
			want: 3,
		},
		{
			name: `text1 = "ace", text2 = "bdf"`,
			args: args{"ace", "bdf"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
