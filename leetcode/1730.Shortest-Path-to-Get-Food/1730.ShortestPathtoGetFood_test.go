package leetcode

import "testing"

func Test_getFood(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `[["X","X","X","X","X"],["X","*","X","O","X"],["X","O","X","#","X"],["X","X","X","X","X"]]`,
			args: args{[][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'X', 'O', 'X'}, {'X', 'O', 'X', '#', 'X'}, {'X', 'X', 'X', 'X', 'X'}}},
			want: -1,
		},
		{
			name: `[["X","X","X","X","X","X"],["X","*","O","O","O","X"],["X","O","O","#","O","X"],["X","X","X","X","X","X"]]`,
			args: args{[][]byte{{'X', 'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'O', 'O', 'O', 'X'}, {'X', 'O', 'O', '#', 'O', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X'}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFood(tt.args.grid); got != tt.want {
				t.Errorf("getFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
