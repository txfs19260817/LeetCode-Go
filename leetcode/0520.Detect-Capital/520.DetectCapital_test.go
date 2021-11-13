package _520_Detect_Capital

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "leetcode",
			args: args{"leetcode"},
			want: true,
		},
		{
			name: "USA",
			args: args{"USA"},
			want: true,
		},
		{
			name: "Google",
			args: args{"Google"},
			want: true,
		},
		{
			name: "LeetCode",
			args: args{"LeetCode"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
