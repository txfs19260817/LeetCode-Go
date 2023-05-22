package leetcode

import "testing"

func Test_countDistinct(t *testing.T) {
	type args struct {
		nums []int
		k    int
		p    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,3,2,2],2,2",
			args: args{[]int{2, 3, 3, 2, 2}, 2, 2},
			want: 11,
		},
		{
			name: "[1,2,3,4],4,1",
			args: args{[]int{1, 2, 3, 4}, 4, 1},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinct(tt.args.nums, tt.args.k, tt.args.p); got != tt.want {
				t.Errorf("countDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
