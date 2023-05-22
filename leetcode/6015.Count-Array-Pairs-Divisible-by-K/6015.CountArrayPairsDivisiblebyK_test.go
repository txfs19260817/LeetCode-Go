package leetcode

import "testing"

func Test_coutPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "[1,2,3,4,5], 2",
			args: args{[]int{1, 2, 3, 4, 5}, 2},
			want: 7,
		},
		{
			name: "[3,2,6,1,8,4,1], 3",
			args: args{[]int{3, 2, 6, 1, 8, 4, 1}, 3},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coutPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("coutPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
