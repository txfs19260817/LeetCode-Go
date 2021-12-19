package _959_Minimum_Operations_to_Make_the_Array_K_Increasing

import "testing"

func Test_kIncreasing(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[12,6,12,6,14,2,13,17,3,8,11,7,4,11,18,8,8,3],1",
			args: args{[]int{12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3}, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kIncreasing(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
