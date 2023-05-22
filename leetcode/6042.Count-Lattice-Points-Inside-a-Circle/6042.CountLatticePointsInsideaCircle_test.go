package leetcode

import "testing"

func Test_countLatticePoints(t *testing.T) {
	type args struct {
		circles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{{2,2,2},{3,4,1}}",
			args: args{[][]int{{2, 2, 2}, {3, 4, 1}}},
			want: 16,
		}, {
			name: "{{8,9,6},{9,8,4}...}",
			args: args{[][]int{{8, 9, 6}, {9, 8, 4}, {4, 1, 1}, {8, 5, 1}, {7, 1, 1}, {6, 7, 5}, {7, 1, 1}, {7, 1, 1}, {5, 5, 3}}},
			want: 141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLatticePoints(tt.args.circles); got != tt.want {
				t.Errorf("countLatticePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
