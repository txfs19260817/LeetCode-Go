package leetcode

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,2,1]",
			args: args{[]int{1, 2, 1}},
			want: 3,
		},
		{
			name: "[0,1,2,2]",
			args: args{[]int{0, 1, 2, 2}},
			want: 3,
		},
		{
			name: "[1,2,3,2,2]",
			args: args{[]int{1, 2, 3, 2, 2}},
			want: 4,
		},
		{
			name: "[3,3,3,1,2,1,1,2,3,3,4]",
			args: args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.tree); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
