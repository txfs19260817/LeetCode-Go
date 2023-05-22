package leetcode

import "testing"

func Test_numOfMinutes(t *testing.T) {
	type args struct {
		n          int
		headID     int
		manager    []int
		informTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 1, headID = 0, manager = [-1], informTime = [0]",
			args: args{1, 0, []int{-1}, []int{0}},
			want: 0,
		},
		{
			name: "n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]",
			args: args{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}},
			want: 1,
		},
		{
			name: "n = 7, headID = 6, manager = [1,2,3,4,5,6,-1], informTime = [0,6,5,4,3,2,1]",
			args: args{7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfMinutes(tt.args.n, tt.args.headID, tt.args.manager, tt.args.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
