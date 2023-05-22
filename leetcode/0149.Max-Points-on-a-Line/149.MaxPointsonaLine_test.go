package leetcode

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[1,1],[2,2],[3,3]]",
			args: args{[][]int{{1, 1}, {2, 2}, {3, 3}}},
			want: 3,
		},
		{
			name: "[[1,1]]",
			args: args{[][]int{{1, 1}}},
			want: 1,
		},
		{
			name: "[[4,5],[4,-1],[4,0]]",
			args: args{[][]int{{4, 5}, {4, -1}, {4, 0}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
