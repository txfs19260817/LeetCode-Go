package leetcode

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
			if got := maxEvents2(tt.args.events); got != tt.want {
				t.Errorf("maxEvents2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxEvents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}})
	}
}

func Benchmark_maxEvents2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxEvents2([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}})
	}
}
