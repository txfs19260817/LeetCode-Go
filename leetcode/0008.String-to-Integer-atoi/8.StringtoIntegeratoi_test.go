package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "   -42"`,
			args: args{"   -42"},
			want: -42,
		},
		{
			name: `s = "4193 with words"`,
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: `s = "words and 987"`,
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: `s = "-91283472332"`,
			args: args{"-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
