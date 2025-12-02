package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:       4,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
				src:     0,
				dst:     3,
				k:       1,
			},
			want: 700,
		},
		{
			name: "example 2",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       1,
			},
			want: 200,
		},
		{
			name: "example 3",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
				src:     0,
				dst:     2,
				k:       0,
			},
			want: 500,
		},
		{
			name: "same source and destination",
			args: args{
				n:       3,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}},
				src:     0,
				dst:     0,
				k:       0,
			},
			want: 0,
		},
		{
			name: "multiple paths with different costs",
			args: args{
				n:       5,
				flights: [][]int{{0, 1, 10}, {0, 2, 20}, {1, 3, 30}, {2, 3, 5}, {3, 4, 10}},
				src:     0,
				dst:     4,
				k:       2,
			},
			want: 35,
		},
		{
			name: "insufficient stops",
			args: args{
				n:       4,
				flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}},
				src:     0,
				dst:     3,
				k:       1,
			},
			want: -1,
		},
		{
			name: "single flight",
			args: args{
				n:       2,
				flights: [][]int{{0, 1, 50}},
				src:     0,
				dst:     1,
				k:       0,
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k))
		})
	}
}
