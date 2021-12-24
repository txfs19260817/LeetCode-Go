package _215_Kth_Largest_Element_in_an_Array

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [3,2,1,5,6,4], k = 2",
			args: args{[]int{3, 2, 1, 5, 6, 4}, 2},
			want: 5,
		},
		{
			name: "nums = [3,2,3,1,2,4,5,5,6], k = 4",
			args: args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
