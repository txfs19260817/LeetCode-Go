package _332_Remove_Palindromic_Subsequences

import "testing"

func Test_removePalindromeSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "baabb",
			args: args{"baabb"},
			want: 2,
		},
		{
			name: "ababa",
			args: args{"ababa"},
			want: 1,
		},
		{
			name: "",
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePalindromeSub(tt.args.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
