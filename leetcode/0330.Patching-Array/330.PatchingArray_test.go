package _330_Patching_Array

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [1,3], n = 6",
			args: args{[]int{1, 3}, 6},
			want: 1,
		},
		{
			name: "nums = [1,2,2], n = 5",
			args: args{[]int{1, 2, 2}, 5},
			want: 0,
		},
		{
			name: "nums = [1,2,31,33], n = 2147483647",
			args: args{[]int{1, 2, 31, 33}, 2147483647},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
