package leetcode

import "testing"

func Test_countElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [-3,3,3,90]",
			args: args{[]int{-3, 3, 3, 90}},
			want: 2,
		},
		{
			name: "nums = [11,7,2,2,15]",
			args: args{[]int{11, 7, 2, 2, 15}},
			want: 2,
		},
		{
			name: "nums = [1]",
			args: args{[]int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.nums); got != tt.want {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
