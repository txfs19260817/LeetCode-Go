package leetcode

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2/3",
			args: args{2, 3},
			want: 0,
		},
		{
			name: "1024/3",
			args: args{1024, 3},
			want: 341,
		},
		{
			name: "1024/-3",
			args: args{1024, -3},
			want: -341,
		},
		{
			name: "-1024/-1",
			args: args{-1024, -1},
			want: 1024,
		},
		{
			name: "2147483647/1",
			args: args{2147483647, 1},
			want: 2147483647,
		},
		{
			name: "2147483648/1",
			args: args{2147483648, 1},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
