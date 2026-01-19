package uber

import "testing"

func TestEarliestAllConnected(t *testing.T) {
	tests := []struct {
		name string
		n    int
		logs []RiderLog
		want int
	}{
		{
			name: "Example",
			n:    4,
			logs: []RiderLog{
				{Time: 20190101, A: 0, B: 1},
				{Time: 20190104, A: 3, B: 2},
				{Time: 20190107, A: 2, B: 0},
				{Time: 20190211, A: 1, B: 2},
				{Time: 20190224, A: 2, B: 3},
			},
			want: 20190107,
		},
		{
			name: "Never connected",
			n:    3,
			logs: []RiderLog{
				{Time: 1, A: 0, B: 1},
			},
			want: -1,
		},
		{
			name: "Already single rider",
			n:    1,
			logs: []RiderLog{},
			want: 0,
		},
		{
			name: "Out of order logs",
			n:    3,
			logs: []RiderLog{
				{Time: 5, A: 0, B: 2},
				{Time: 1, A: 0, B: 1},
			},
			want: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EarliestAllConnected(tc.n, tc.logs); got != tc.want {
				t.Fatalf("EarliestAllConnected() = %d, want %d", got, tc.want)
			}
		})
	}
}
