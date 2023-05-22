package leetcode

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [3,2,3]",
			args: args{[]int{3, 2, 3}},
			want: []int{3},
		},
		{
			name: "nums = [3]",
			args: args{[]int{3}},
			want: []int{3},
		},
		{
			name: "nums = [1,3]",
			args: args{[]int{1, 3}},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
