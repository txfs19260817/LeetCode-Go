package leetcode

import "testing"

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
		time      []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1: Basic dependency chain",
			args: args{
				n:         3,
				relations: [][]int{{1, 3}, {2, 3}},
				time:      []int{3, 2, 5},
			},
			want: 8,
		},
		{
			name: "Example 2: Linear dependency",
			args: args{
				n:         5,
				relations: [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}},
				time:      []int{1, 2, 3, 4, 5},
			},
			want: 12,
		},
		{
			name: "Single course",
			args: args{
				n:         1,
				relations: [][]int{},
				time:      []int{10},
			},
			want: 10,
		},
		{
			name: "No dependencies - all parallel",
			args: args{
				n:         4,
				relations: [][]int{},
				time:      []int{5, 3, 7, 2},
			},
			want: 7,
		},
		{
			name: "Long chain dependency",
			args: args{
				n:         4,
				relations: [][]int{{1, 2}, {2, 3}, {3, 4}},
				time:      []int{1, 2, 3, 4},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime(tt.args.n, tt.args.relations, tt.args.time); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
