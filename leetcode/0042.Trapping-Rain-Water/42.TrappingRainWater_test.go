package leetcode

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "height = [4,2,0,3,2,5]",
			args: args{[]int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
		{
			name: "height = [1,1,1]",
			args: args{[]int{1, 1, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
