package leetcode

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				arr: []int{3, 0, 1, 1, 9, 7},
				a:   7,
				b:   2,
				c:   3,
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				arr: []int{1, 1, 2, 2, 3},
				a:   0,
				b:   0,
				c:   1,
			},
			want: 0,
		},
		{
			name: "custom case",
			args: args{
				arr: []int{7, 3, 7, 3, 12, 1, 12, 2, 3},
				a:   5,
				b:   8,
				c:   1,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
