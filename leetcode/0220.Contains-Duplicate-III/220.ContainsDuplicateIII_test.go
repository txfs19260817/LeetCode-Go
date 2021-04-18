package _220_Contains_Duplicate_III

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [1,5,9,1,5,9], k = 2, t = 3",
			args: args{[]int{1, 5, 9, 1, 5, 9}, 2, 3},
			want: false,
		},
		{
			name: "nums = [1,2,3,1], k = 3, t = 0",
			args: args{[]int{1, 2, 3, 1}, 3, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
