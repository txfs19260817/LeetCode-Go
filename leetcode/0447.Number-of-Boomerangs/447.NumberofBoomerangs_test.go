package _447_Number_of_Boomerangs

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "points = [[0,0],[1,0],[2,0]]",
			args: args{[][]int{{0, 0}, {1, 0}, {2, 0}}},
			want: 2,
		},
		{
			name: "points = [[1,1]]",
			args: args{[][]int{{1, 1}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
