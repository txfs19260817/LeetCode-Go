package _174_Dungeon_Game

import "testing"

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]",
			args: args{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}},
			want: 7,
		},
		{
			name: "dungeon = [[0]]",
			args: args{[][]int{{0}}},
			want: 1,
		},
		{
			name: "dungeon = [[1,-3,3],[0,-2,0],[-3,-3,-3]]",
			args: args{[][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
