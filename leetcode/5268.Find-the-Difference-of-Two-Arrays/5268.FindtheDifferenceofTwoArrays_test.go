package _268_Find_the_Difference_of_Two_Arrays

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums1 = [1,2,3,3], nums2 = [1,1,2,2]",
			args: args{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}},
			want: [][]int{{3}, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
