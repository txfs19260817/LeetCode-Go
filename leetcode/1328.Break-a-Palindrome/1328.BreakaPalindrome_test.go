package _328_Break_a_Palindrome

import "testing"

func Test_breakPalindrome(t *testing.T) {
	type args struct {
		palindrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abccba",
			args: args{"abccba"},
			want: "aaccba",
		},
		{
			name: "a",
			args: args{"a"},
			want: "",
		},
		{
			name: "aa",
			args: args{"aa"},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakPalindrome(tt.args.palindrome); got != tt.want {
				t.Errorf("breakPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
