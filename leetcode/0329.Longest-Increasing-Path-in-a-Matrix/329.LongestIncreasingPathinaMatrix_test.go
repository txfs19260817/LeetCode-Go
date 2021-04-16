package _329_Longest_Increasing_Path_in_a_Matrix

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "matrix = [[9,9,4],[6,6,8],[2,1,1]]",
			args: args{[][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
