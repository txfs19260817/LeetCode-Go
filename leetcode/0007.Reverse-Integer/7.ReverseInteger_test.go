package _007_Reverse_Integer

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "x = 123",
			args: args{123},
			want: 321,
		},
		{
			name: "x = -123",
			args: args{-123},
			want: -321,
		},
		{
			name: "x = 120",
			args: args{120},
			want: 21,
		},
		{
			name: "x = 0",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
