package leetcode

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "intervals = [[0,30],[5,10],[15,20]]",
			args: args{[][]int{{0, 30}, {5, 10}, {15, 20}}},
			want: 2,
		},
		{
			name: "intervals = [[7,10],[2,4]]",
			args: args{[][]int{{7, 10}, {2, 4}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
