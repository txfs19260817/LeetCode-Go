package _237_Count_Positions_on_Street_With_Required_Brightness

import "testing"

func Test_meetRequirement(t *testing.T) {
	type args struct {
		n           int
		lights      [][]int
		requirement []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 5, lights = [[0,1],[2,1],[3,2]], requirement = [0,2,1,4,1]",
			args: args{5, [][]int{{0, 1}, {2, 1}, {3, 2}}, []int{0, 2, 1, 4, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := meetRequirement(tt.args.n, tt.args.lights, tt.args.requirement); got != tt.want {
				t.Errorf("meetRequirement() = %v, want %v", got, tt.want)
			}
		})
	}
}
