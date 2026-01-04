package leetcode

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{n: "123"},
			want: "121",
		},
		{
			name: "Example 2",
			args: args{n: "1"},
			want: "0",
		},
		{
			name: "Small number 10",
			args: args{n: "10"},
			want: "9",
		},
		{
			name: "Small number 11",
			args: args{n: "11"},
			want: "9",
		},
		{
			name: "Number 100",
			args: args{n: "100"},
			want: "99",
		},
		{
			name: "Number 1283",
			args: args{n: "1283"},
			want: "1331",
		},
		{
			name: "Already Palindrome 99",
			args: args{n: "99"},
			want: "101",
		},
		{
			name: "Large number",
			args: args{n: "807045053224792883"},
			want: "807045053350540708",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
