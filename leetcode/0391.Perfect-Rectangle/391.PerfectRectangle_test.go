package leetcode

import "testing"

func Test_isRectangleCover(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]",
			args: args{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}},
			want: true,
		},
		{
			name: "rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]",
			args: args{[][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}},
			want: false,
		},
		{
			name: "rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]",
			args: args{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}}},
			want: false,
		},
		{
			name: "rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]",
			args: args{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleCover(tt.args.rectangles); got != tt.want {
				t.Errorf("isRectangleCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
