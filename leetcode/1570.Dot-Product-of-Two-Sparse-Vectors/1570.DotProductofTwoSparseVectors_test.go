package leetcode

import (
	"reflect"
	"testing"
)

func TestSparseVector(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]",
			args: args{[]int{1, 0, 0, 2, 3}, []int{0, 3, 0, 4, 0}},
			want: 8,
		},
		{
			name: "nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]",
			args: args{[]int{0, 1, 0, 0, 0}, []int{0, 0, 0, 0, 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := Constructor(tt.args.nums1)
			v2 := Constructor(tt.args.nums2)
			if got := v1.dotProduct(v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
			if got := v2.dotProduct(v1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
