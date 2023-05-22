package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [4,5,6,7,0,1,2], target = 0",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			want: 4,
		},
		{
			name: "nums = [4,5,6,7,0,1,2], target = 1",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 1},
			want: 5,
		},
		{
			name: "nums = [4,5,6,7], target = 7",
			args: args{[]int{4, 5, 6, 7}, 7},
			want: 3,
		},
		{
			name: "nums = [4,5,6,7], target = 4",
			args: args{[]int{4, 5, 6, 7}, 4},
			want: 0,
		},
		{
			name: "nums = [4,5,6,7,0,1,2], target = 3",
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
