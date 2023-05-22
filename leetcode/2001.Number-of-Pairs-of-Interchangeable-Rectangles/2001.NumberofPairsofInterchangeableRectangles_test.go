package leetcode

import "testing"

func Test_interchangeableRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "rectangles = [[4,8],[3,6],[10,20],[15,30]]",
			args: args{[][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}},
			want: 6,
		},
		{
			name: "rectangles = [[4,5],[7,8]]",
			args: args{[][]int{{4, 5}, {7, 8}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interchangeableRectangles(tt.args.rectangles); got != tt.want {
				t.Errorf("interchangeableRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
