package leetcode

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "costs = [[17,2,17],[16,16,5],[14,3,19]]",
			args: args{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}},
			want: 10,
		},
		{
			name: "costs = [[7,6,2]]",
			args: args{[][]int{{7, 6, 2}}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.costs); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
