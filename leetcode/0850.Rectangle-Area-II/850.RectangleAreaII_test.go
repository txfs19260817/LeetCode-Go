package leetcode

import "testing"

func Test_rectangleArea(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "rectangles = [[0,0,2,2],[1,0,2,3],[1,0,3,1]]",
			args: args{[][]int{{0, 0, 2, 2}, {1, 0, 2, 3}, {1, 0, 3, 1}}},
			want: 6,
		},
		{
			name: "rectangles = [[0,0,1000000000,1000000000]]",
			args: args{[][]int{{0, 0, 1000000000, 1000000000}}},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rectangleArea(tt.args.rectangles); got != tt.want {
				t.Errorf("rectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
