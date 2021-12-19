package _353_Maximum_Number_of_Events_That_Can_Be_Attended

import "testing"

func Test_maxEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "events = [[1,2],[2,3],[3,4]]",
			args: args{[][]int{{1, 2}, {2, 3}, {3, 4}}},
			want: 3,
		}, {
			name: "events = [[1,2],[1,2],[3,3],[1,5],[1,5]]",
			args: args{[][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
