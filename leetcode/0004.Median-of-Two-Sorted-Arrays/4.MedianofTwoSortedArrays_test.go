package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "nums1 = [1,3], nums2 = [2]",
			args: args{[]int{1, 3}, []int{2}},
			want: 2.0,
		},
		{
			name: "nums1 = [1,3], nums2 = [2,4]",
			args: args{[]int{1, 3}, []int{2, 4}},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
