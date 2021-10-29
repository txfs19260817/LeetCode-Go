package _997_Find_the_Town_Judge

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 2, trust = [[1,2]]",
			args: args{2, [][]int{{1, 2}}},
			want: 2,
		},
		{
			name: "n = 3, trust = [[1,3],[2,3],[3,1]]",
			args: args{3, [][]int{{1, 3}, {2, 3}, {3, 1}}},
			want: -1,
		},
		{
			name: "n = 1, trust = []",
			args: args{1, [][]int{}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
