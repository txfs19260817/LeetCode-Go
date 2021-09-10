package _764_Largest_Plus_Sign

import "testing"

func Test_orderOfLargestPlusSign(t *testing.T) {
	type args struct {
		n     int
		mines [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 5, mines = [[4,2]]",
			args: args{5, [][]int{{4, 2}}},
			want: 2,
		},
		{
			name: "n = 1, mines = [[0,0]]",
			args: args{1, [][]int{{0, 0}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.args.n, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargestPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
