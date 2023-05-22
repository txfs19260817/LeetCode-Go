package leetcode

import "testing"

func Test_minimalKSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "nums = [5,6], k = 6",
			args: args{[]int{5, 6}, 6},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalKSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimalKSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
