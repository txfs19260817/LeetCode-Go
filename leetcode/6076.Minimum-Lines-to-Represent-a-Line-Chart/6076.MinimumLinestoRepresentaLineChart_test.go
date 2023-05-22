package leetcode

import "testing"

func Test_minimumLines(t *testing.T) {
	type args struct {
		stockPrices [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]",
			args: args{[][]int{{1, 7}, {2, 6}, {3, 5}, {4, 4}, {5, 4}, {6, 3}, {7, 2}, {8, 1}}},
			want: 3,
		},
		{
			name: "stockPrices = [[1,7],[3,7]]",
			args: args{[][]int{{1, 7}, {3, 7}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumLines(tt.args.stockPrices); got != tt.want {
				t.Errorf("minimumLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
