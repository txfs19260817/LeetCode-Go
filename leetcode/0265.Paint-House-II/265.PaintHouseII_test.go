package _265_Paint_House_II

import "testing"

func Test_minCostII(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "costs = [[1,5,3],[2,9,4]]",
			args: args{[][]int{{1,5,3}, {2,9,4}}},
			want: 5,
		},
		{
			name: "costs = [[1,3]]",
			args: args{[][]int{{1,3}}},
			want: 1,
		},
		{
			name: "costs = [[20,19,11,13,12,16,16,17,15,9,5,18],[3,8,15,17,19,8,18,3,11,6,7,12],[15,4,11,1,18,2,10,9,3,6,4,15]]",
			args: args{[][]int{{20,19,11,13,12,16,16,17,15,9,5,18},{3,8,15,17,19,8,18,3,11,6,7,12},{15,4,11,1,18,2,10,9,3,6,4,15}}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostII(tt.args.costs); got != tt.want {
				t.Errorf("minCostII() = %v, want %v", got, tt.want)
			}
		})
	}
}
