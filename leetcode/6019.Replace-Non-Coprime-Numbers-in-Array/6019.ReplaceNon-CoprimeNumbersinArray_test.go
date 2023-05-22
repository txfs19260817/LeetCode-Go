package leetcode

import (
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[2,3,...,2,3,6,6]",
			args: args{[]int{2, 3, 2, 3, 6, 6}},
			want: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceNonCoprimes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
