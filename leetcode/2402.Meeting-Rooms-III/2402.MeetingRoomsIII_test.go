package leetcode

import "testing"

func Test_mostBooked(t *testing.T) {
	type args struct {
		n        int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:        2,
				meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				n:        3,
				meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			},
			want: 1,
		},
		{
			name: "single room delays every overlap",
			args: args{
				n:        1,
				meetings: [][]int{{0, 4}, {1, 3}, {2, 5}, {6, 8}},
			},
			want: 0,
		},
		{
			name: "meetings are sorted by start time",
			args: args{
				n:        2,
				meetings: [][]int{{5, 6}, {0, 10}, {1, 2}, {2, 3}},
			},
			want: 1,
		},
		{
			name: "lowest idle room is reused first",
			args: args{
				n:        3,
				meetings: [][]int{{0, 2}, {1, 3}, {2, 4}, {4, 6}, {5, 7}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostBooked(tt.args.n, tt.args.meetings); got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
