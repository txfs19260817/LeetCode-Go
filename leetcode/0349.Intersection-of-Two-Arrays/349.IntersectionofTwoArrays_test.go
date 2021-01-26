package _349_Intersection_of_Two_Arrays

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums1 = [1,2,2,1], nums2 = [2,2]",
			args: args{[]int{1,2,2,1}, []int{2,2}},
			want: []int{2},
		},
		{
			name: "nums1 = [9,5], nums2 = [9,4,9,8,4]",
			args: args{[]int{9,5}, []int{9,4,9,8,4}},
			want: []int{9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_intersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intersection([]int{9,5}, []int{9,4,9,8,4})
	}
}
