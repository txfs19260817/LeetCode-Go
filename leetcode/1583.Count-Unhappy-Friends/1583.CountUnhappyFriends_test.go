package _583_Count_Unhappy_Friends

import "testing"

func Test_unhappyFriends(t *testing.T) {
	type args struct {
		n           int
		preferences [][]int
		pairs       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 4, preferences = [[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]], pairs = [[0, 1], [2, 3]]",
			args: args{4, [][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}}, [][]int{{0, 1}, {2, 3}}},
			want: 2,
		},
		{
			name: "n = 2, preferences = [[1], [0]], pairs = [[1, 0]]",
			args: args{2, [][]int{{1}, {0}}, [][]int{{1, 0}}},
			want: 0,
		},
		{
			name: "n = 4, preferences = [[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]], pairs = [[1, 3], [0, 2]]",
			args: args{4, [][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}}, [][]int{{1, 3}, {0, 2}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unhappyFriends(tt.args.n, tt.args.preferences, tt.args.pairs); got != tt.want {
				t.Errorf("unhappyFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
