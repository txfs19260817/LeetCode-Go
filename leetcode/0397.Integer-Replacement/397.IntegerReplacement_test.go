package leetcode

import (
	"math"
	"testing"
)

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1: 8 -> 4 -> 2 -> 1",
			args: args{n: 8},
			want: 3,
		},
		{
			name: "example 2: 7 -> 8 -> 4 -> 2 -> 1",
			args: args{n: 7},
			want: 4,
		},
		{
			name: "example 3: 4 -> 2 -> 1",
			args: args{n: 4},
			want: 2,
		},
		{
			name: "n = 1",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "n = 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "n = 3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "n = 15",
			args: args{n: 15},
			want: 5,
		},
		{
			name: "n = 27",
			args: args{n: 27},
			want: 7,
		},
		{
			name: "n = 100",
			args: args{n: 100},
			want: 8,
		},
		{
			name: "n = MaxInt32",
			args: args{n: math.MaxInt32},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
