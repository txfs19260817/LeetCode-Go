package leetcode

import "testing"

func Test_minSideJumps(t *testing.T) {
	type args struct {
		obstacles []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{obstacles: []int{0, 1, 2, 3, 0}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{obstacles: []int{0, 1, 1, 3, 3, 0}},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{obstacles: []int{0, 2, 1, 0, 3, 0}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSideJumps(tt.args.obstacles); got != tt.want {
				t.Errorf("minSideJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
