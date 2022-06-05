package _091_Partition_Array_Such_That_Maximum_Difference_Is_K

import "testing"

func Test_partitionArray(t *testing.T) {
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
			name: "[2,2,4,5],0",
			args: args{[]int{2, 2, 4, 5}, 0},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("partitionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
