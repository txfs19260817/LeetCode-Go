package _645_Set_Mismatch

import (
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,2,2,4]",
			args: args{[]int{1, 2, 2, 4}},
			want: []int{2, 3},
		},
		{
			name: "nums = [1,1]",
			args: args{[]int{1, 1}},
			want: []int{1, 2},
		},
		{
			name: "nums = [3,2,3,4,6,5]",
			args: args{[]int{3, 2, 3, 4, 6, 5}},
			want: []int{3, 1},
		},
		{
			name: "nums = [3,2,2]",
			args: args{[]int{3, 2, 2}},
			want: []int{2, 1},
		},
		{
			name: "nums = [8,7,3,5,3,6,1,4]",
			args: args{[]int{8, 7, 3, 5, 3, 6, 1, 4}},
			want: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findErrorNums(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
