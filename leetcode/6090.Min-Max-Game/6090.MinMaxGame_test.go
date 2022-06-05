package _090_Min_Max_Game

import "testing"

func Test_minMaxGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[70,38,21,22]",
			args: args{[]int{70, 38, 21, 22}},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxGame(tt.args.nums); got != tt.want {
				t.Errorf("minMaxGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
