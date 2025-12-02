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
		{
			name: "all meetings overlap - maximum rooms needed",
			args: args{[][]int{{0, 5}, {1, 6}, {2, 7}, {3, 8}, {4, 9}}},
			want: 5,
		},
		{
			name: "meetings end exactly when another starts - reuse rooms",
			args: args{[][]int{{0, 5}, {5, 10}, {10, 15}, {15, 20}}},
			want: 1,
		},
		{
			name: "single meeting",
			args: args{[][]int{{1, 5}}},
			want: 1,
		},
		{
			name: "empty intervals",
			args: args{[][]int{}},
			want: 0,
		},
		{
			name: "nested meetings - all overlap",
			args: args{[][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}, {5, 6}}},
			want: 5,
		},
		{
			name: "complex overlapping pattern",
			args: args{[][]int{{0, 10}, {5, 15}, {10, 20}, {15, 25}, {20, 30}}},
			want: 2,
		},
		{
			name: "multiple peaks requiring different rooms",
			args: args{[][]int{{0, 5}, {1, 3}, {2, 4}, {6, 10}, {7, 9}, {8, 11}, {12, 15}, {13, 14}}},
			want: 3,
		},
		{
			name: "all meetings at same time",
			args: args{[][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}}},
			want: 5,
		},
		{
			name: "large number of non-overlapping meetings",
			args: args{[][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}, {12, 13}, {14, 15}}},
			want: 1,
		},
		{
			name: "meeting completely contained in another",
			args: args{[][]int{{0, 20}, {5, 10}}},
			want: 2,
		},
		{
			name: "alternating pattern",
			args: args{[][]int{{0, 5}, {10, 15}, {5, 10}, {15, 20}}},
			want: 1,
		},
		{
			name: "three overlapping groups",
			args: args{[][]int{{0, 3}, {1, 4}, {2, 5}, {10, 13}, {11, 14}, {12, 15}, {20, 23}, {21, 24}, {22, 25}}},
			want: 3,
		},
		{
			name: "single point meetings",
			args: args{[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
			want: 1,
		},
		{
			name: "very long meeting with many short overlaps",
			args: args{[][]int{{0, 100}, {10, 20}, {30, 40}, {50, 60}, {70, 80}}},
			want: 2,
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
