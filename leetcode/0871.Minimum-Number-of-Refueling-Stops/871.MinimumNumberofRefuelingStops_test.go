package leetcode

import "testing"

func Test_minRefuelStops(t *testing.T) {
	type args struct {
		target    int
		startFuel int
		stations  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "target = 1, startFuel = 1, stations = []",
			args: args{1, 1, [][]int{}},
			want: 0,
		},
		{
			name: "target = 100, startFuel = 1, stations = [[10,100]]",
			args: args{100, 1, [][]int{{10, 100}}},
			want: -1,
		},
		{
			name: "target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]",
			args: args{100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}},
			want: 2,
		},
		{
			name: "target = 100, startFuel = 50, stations = [[25,25],[50,50]]",
			args: args{100, 50, [][]int{{25, 25}, {50, 50}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRefuelStops(tt.args.target, tt.args.startFuel, tt.args.stations); got != tt.want {
				t.Errorf("minRefuelStops() = %v, want %v", got, tt.want)
			}
			if got := minRefuelStops2(tt.args.target, tt.args.startFuel, tt.args.stations); got != tt.want {
				t.Errorf("minRefuelStops2() = %v, want %v", got, tt.want)
			}
		})
	}
}
