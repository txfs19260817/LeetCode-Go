package _317_Shortest_Distance_from_All_Buildings

import "testing"

var large = func() (m [][]int) {
	m = append(m, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	for i := 0; i < 48; i++ {
		m = append(m, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	}
	m = append(m, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	return
}()

func Test_shortestDistance(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "grid = [[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]",
			args: args{[][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}},
			want: 7,
		},
		{
			name: "grid = [[1,0]]",
			args: args{[][]int{{1, 0}}},
			want: 1,
		},
		{
			name: "grid = [[1]]",
			args: args{[][]int{{1}}},
			want: -1,
		},
		{
			name: "grid = [[1,2,0]]",
			args: args{[][]int{{1, 2, 0}}},
			want: -1,
		},
		{
			name: "grid = [[1,1,1,1,1,0]...",
			args: args{[][]int{{1, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0, 1}, {0, 1, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 1}, {1, 0, 1, 0, 0, 1}, {1, 0, 0, 0, 0, 1}, {0, 1, 1, 1, 1, 0}}},
			want: 88,
		},
		{
			name: "large",
			args: args{large},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.grid); got != tt.want {
				t.Errorf("shortestDistance() = %v, want %v", got, tt.want)
			}
			if got := shortestDistance1(tt.args.grid); got != tt.want {
				t.Errorf("shortestDistance1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_shortestDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shortestDistance(large)
	}
}

func Benchmark_shortestDistance1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shortestDistance1(large)
	}
}
