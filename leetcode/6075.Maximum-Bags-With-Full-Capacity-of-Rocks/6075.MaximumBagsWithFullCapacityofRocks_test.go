package leetcode

import "testing"

func Test_maximumBags(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2",
			args: args{[]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBags(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks); got != tt.want {
				t.Errorf("maximumBags() = %v, want %v", got, tt.want)
			}
		})
	}
}
