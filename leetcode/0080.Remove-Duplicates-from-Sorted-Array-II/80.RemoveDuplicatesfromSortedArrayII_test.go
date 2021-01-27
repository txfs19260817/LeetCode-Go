package _080_Remove_Duplicates_from_Sorted_Array_II

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,1,1,2,2,3]",
			args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: 5,
		},
		{
			name: "nums = [0,0,1,1,1,1,2,3,3]",
			args: args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
