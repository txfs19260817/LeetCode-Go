package leetcode

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "intervals = [[1,4],[3,6],[2,8]]",
			args: args{[][]int{{1, 4}, {3, 6}, {2, 8}}},
			want: 2,
		},
		{
			name: "intervals = [[1,4],[1,3],[3,6],[2,8]]",
			args: args{[][]int{{1, 4}, {1, 3}, {3, 6}, {2, 8}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
