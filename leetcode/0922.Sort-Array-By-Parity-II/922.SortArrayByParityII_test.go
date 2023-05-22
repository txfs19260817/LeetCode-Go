package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [4,2,5,7]",
			args: args{[]int{4, 2, 5, 7}},
			want: []int{4, 5, 2, 7},
		},
		{
			name: "nums = [2,5]",
			args: args{[]int{2, 5}},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
			if got := sortArrayByParityII2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII2() = %v, want %v", got, tt.want)
			}
			if got := sortArrayByParityII3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII3() = %v, want %v", got, tt.want)
			}
		})
	}
}
