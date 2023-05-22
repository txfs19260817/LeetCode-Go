package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "race a car",
			args: args{"race a car"},
			want: false,
		},
		{
			name: "0P",
			args: args{"0P"},
			want: false,
		},
		{
			name: "a",
			args: args{"a"},
			want: true,
		},
		{
			name: "",
			args: args{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
