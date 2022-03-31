package _236_Minimum_Deletions_to_Make_Array_Beautiful

import "testing"

func Test_minDeletion(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,1,2,2,3,3]",
			args: args{[]int{1, 1, 2, 2, 3, 3}},
			want: 2,
		},
		{
			name: "nums = [1,6,0,9,8,...]",
			args: args{[]int{1, 6, 0, 9, 8, 8, 4, 1, 7, 1, 1, 8, 9, 1, 9, 1, 2, 3, 7, 6, 6, 8, 7, 5, 9, 7, 3, 0, 4, 9, 1, 8, 3, 3, 2, 4, 2, 6, 2, 8, 9, 2, 8, 4, 0, 1, 0, 9, 9, 5}},
			want: 4,
		},
		{
			name: "nums = [5,1,5,4,8,1,4,4,1,9,2,2,2,5,1]",
			args: args{[]int{5, 1, 5, 4, 8, 1, 4, 4, 1, 9, 2, 2, 2, 5, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletion(tt.args.nums); got != tt.want {
				t.Errorf("minDeletion() = %v, want %v", got, tt.want)
			}
		})
	}
}
