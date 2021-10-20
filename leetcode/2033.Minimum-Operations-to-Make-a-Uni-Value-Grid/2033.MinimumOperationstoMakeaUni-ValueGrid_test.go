package _033_Minimum_Operations_to_Make_a_Uni_Value_Grid

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		grid [][]int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[2,4],[6,8]],2",
			args: args{[][]int{{2, 4}, {6, 8}}, 2},
			want: 4,
		},
		{
			name: "[[1,5],[2,3]],1",
			args: args{[][]int{{1, 5}, {2, 3}}, 1},
			want: 5,
		},
		{
			name: "[[1,2],[3,4]],2",
			args: args{[][]int{{1, 2}, {3, 4}}, 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.grid, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
