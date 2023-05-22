package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		K      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "points = [[3,3],[5,-1],[-2,4]], k = 2",
			args: args{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2},
			want: [][]int{{3, 3}, {-2, 4}},
		},
		{
			name: "points = [[1,3],[-2,2]], k = 1",
			args: args{[][]int{{1, 3}, {-2, 2}}, 1},
			want: [][]int{{-2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, kClosest(tt.args.points, tt.args.K))
			assert.ElementsMatch(t, tt.want, kClosest2(tt.args.points, tt.args.K))
		})
	}
}

func Benchmark_kClosest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
		kClosest([][]int{{3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}}, 1)
	}
}

func Benchmark_kClosest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosest2([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
		kClosest2([][]int{{3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}, {3, 3}}, 1)
	}
}
