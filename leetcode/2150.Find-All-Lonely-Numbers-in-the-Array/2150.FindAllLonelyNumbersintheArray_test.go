package leetcode

import (
	"reflect"
	"testing"
)

func Test_findLonely(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [10,6,5,8]",
			args: args{[]int{10, 6, 5, 8}},
			want: []int{10, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLonely(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLonely() = %v, want %v", got, tt.want)
			}
		})
	}
}
