package _905_Second_Minimum_Time_to_Reach_Destination

import "testing"

func Test_secondMinimum(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		time   int
		change int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 5, edges = [[1,2],[1,3],[1,4],[3,4],[4,5]], time = 3, change = 5",
			args: args{5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5},
			want: 13,
		},
		{
			name: "n = 2, edges = [[1,2]], time = 3, change = 2",
			args: args{2, [][]int{{1, 2}}, 3, 2},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondMinimum(tt.args.n, tt.args.edges, tt.args.time, tt.args.change); got != tt.want {
				t.Errorf("secondMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
