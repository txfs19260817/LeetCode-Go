package _875_Koko_Eating_Bananas

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "piles = [3,6,7,11], h = 8",
			args: args{[]int{3, 6, 7, 11}, 8},
			want: 4,
		},
		{
			name: "piles = [30,11,23,4,20], h = 5",
			args: args{[]int{30, 11, 23, 4, 20}, 5},
			want: 30,
		},
		{
			name: "piles = [30,11,23,4,20], h = 6",
			args: args{[]int{30, 11, 23, 4, 20}, 6},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
