package _966_Recover_the_Original_Array

import (
	"reflect"
	"testing"
)

func Test_recoverArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [2,10,6,4,8,12]",
			args: args{[]int{2, 10, 6, 4, 8, 12}},
			want: []int{3, 7, 11},
		},
		{
			name: "nums = [1,1,3,3]",
			args: args{[]int{1, 1, 3, 3}},
			want: []int{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
