package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "target = 7, nums = [2,3,1,2,4,3]",
			args: args{7, []int{2, 3, 1, 2, 4, 3}},
			want: 2,
		},
		{
			name: "target = 11, nums = [1,2,3,4,5]",
			args: args{11, []int{1, 2, 3, 4, 5}},
			want: 3,
		},
		{
			name: "target = 11, nums = [1,1,1,1,1,1,1,1]",
			args: args{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
			if got := minSubArrayLen2(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen2() = %v, want %v", got, tt.want)
			}
		})
	}
}
