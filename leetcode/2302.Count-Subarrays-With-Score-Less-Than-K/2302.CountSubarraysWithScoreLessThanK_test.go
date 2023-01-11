package _302_Count_Subarrays_With_Score_Less_Than_K

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "nums = [2,1,4,3,5], k = 10",
			args: args{[]int{2, 1, 4, 3, 5}, 10},
			want: 6,
		},
		{
			name: "nums = [1,1,1], k = 5",
			args: args{[]int{1, 1, 1}, 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
