package leetcode

import "testing"

func Test_maxRectangleArea(t *testing.T) {
	type args struct {
		xCoord []int
		yCoord []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example 1",
			args: args{
				xCoord: []int{1, 1, 3, 3},
				yCoord: []int{1, 3, 1, 3},
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				xCoord: []int{1, 1, 3, 3, 2},
				yCoord: []int{1, 3, 1, 3, 2},
			},
			want: -1,
		},
		{
			name: "example 3",
			args: args{
				xCoord: []int{1, 1, 3, 3, 1, 3},
				yCoord: []int{1, 3, 1, 3, 2, 2},
			},
			want: 2,
		},
		{
			name: "no rectangle",
			args: args{
				xCoord: []int{0, 1, 2},
				yCoord: []int{0, 1, 2},
			},
			want: -1,
		},
		{
			name: "border point blocks rectangle",
			args: args{
				xCoord: []int{0, 0, 2, 2, 1},
				yCoord: []int{0, 2, 0, 2, 0},
			},
			want: -1,
		},
		{
			name: "multiple rectangles pick max area",
			args: args{
				xCoord: []int{0, 0, 2, 2, 5, 5, 8, 8},
				yCoord: []int{0, 2, 0, 2, 1, 5, 1, 5},
			},
			want: 12,
		},
		{
			name: "large blocked small valid",
			args: args{
				xCoord: []int{0, 0, 4, 4, 2, 0, 2},
				yCoord: []int{0, 4, 0, 4, 2, 2, 0},
			},
			want: 4,
		},
		{
			name: "negative coordinates",
			args: args{
				xCoord: []int{-3, -3, 1, 1},
				yCoord: []int{-1, 2, -1, 2},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRectangleArea(tt.args.xCoord, tt.args.yCoord); got != tt.want {
				t.Errorf("maxRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
