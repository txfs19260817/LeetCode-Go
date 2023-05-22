package leetcode

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [5,7,7,8,8,10], target = 8",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 8},
			want: []int{3, 4},
		},
		{
			name: "nums = [5,7,7,8,8,10], target = 6",
			args: args{[]int{5, 7, 7, 8, 8, 10}, 6},
			want: []int{-1, -1},
		},
		{
			name: "nums = [], target = 6",
			args: args{[]int{}, 6},
			want: []int{-1, -1},
		},
		{
			name: "nums = [1], target = 6",
			args: args{[]int{1}, 6},
			want: []int{-1, -1},
		},
		{
			name: "nums = [6,6], target = 6",
			args: args{[]int{6, 6}, 6},
			want: []int{0, 1},
		},
		{
			name: "nums = [1,1,2], target = 1",
			args: args{[]int{1, 1, 2}, 1},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
