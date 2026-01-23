package leetcode

import "testing"

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1_transfer_at_stop",
			args: args{
				routes: [][]int{{1, 2, 7}, {3, 6, 7}},
				source: 1,
				target: 6,
			},
			want: 2,
		},
		{
			name: "example2_unreachable",
			args: args{
				routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}},
				source: 15,
				target: 12,
			},
			want: -1,
		},
		{
			name: "single_transfer",
			args: args{
				routes: [][]int{{2}, {2, 8}},
				source: 8,
				target: 2,
			},
			want: 1,
		},
		{
			name: "source_equals_target",
			args: args{
				routes: [][]int{{1, 2, 3}, {4, 5, 6}},
				source: 5,
				target: 5,
			},
			want: 0,
		},
		{
			name: "direct_bus",
			args: args{
				routes: [][]int{{1, 4, 7}, {2, 3, 5}},
				source: 1,
				target: 7,
			},
			want: 1,
		},
		{
			name: "source_not_in_routes",
			args: args{
				routes: [][]int{{1, 2, 3}, {4, 5, 6}},
				source: 9,
				target: 1,
			},
			want: -1,
		},
		{
			name: "three_buses_chain",
			args: args{
				routes: [][]int{{1, 2}, {2, 3}, {3, 4}},
				source: 1,
				target: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
