package leetcode

import "testing"

func Test_connectSticks(t *testing.T) {
	type args struct {
		sticks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sticks = [1,8,3,5]",
			args: args{[]int{1, 8, 3, 5}},
			want: 30,
		},
		{
			name: "sticks = [5]",
			args: args{[]int{5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connectSticks(tt.args.sticks); got != tt.want {
				t.Errorf("connectSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}
