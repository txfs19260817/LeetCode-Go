package _312_Burst_Balloons

import "testing"

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,1,5,8]",
			args: args{[]int{3, 1, 5, 8}},
			want: 167,
		},
		{
			name: "nums = [1,5]",
			args: args{[]int{1, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
