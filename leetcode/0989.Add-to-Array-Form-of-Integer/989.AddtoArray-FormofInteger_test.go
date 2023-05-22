package leetcode

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	type args struct {
		num []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "num = [1,2,0,0], k = 34",
			args: args{[]int{1, 2, 0, 0}, 34},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "num = [2,7,4], k = 181",
			args: args{[]int{2, 7, 4}, 181},
			want: []int{4, 5, 5},
		},
		{
			name: "num = [2,1,5], k = 806",
			args: args{[]int{2, 1, 5}, 806},
			want: []int{1, 0, 2, 1},
		},
		{
			name: "num = [9,9,9,9,9,9,9,9,9,9], k = 1",
			args: args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addToArrayForm(tt.args.num, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
