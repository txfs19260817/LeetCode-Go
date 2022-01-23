package _149_Rearrange_Array_Elements_by_Sign

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[-1,1]",
			args: args{[]int{-1, 1}},
			want: []int{1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
