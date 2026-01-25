package leetcode

import (
	"reflect"
	"testing"
)

func Test_numIslands2(t *testing.T) {
	type args struct {
		m         int
		n         int
		positions [][]int
	}
	tests := []struct {
		name         string
		args         args
		want         []int
		wantMaxSizes []int
	}{
		{
			name: "Example 1",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}},
			},
			want:         []int{1, 1, 2, 3},
			wantMaxSizes: []int{1, 2, 2, 2},
		},
		{
			name: "Example with duplicates",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {0, 0}},
			},
			want:         []int{1, 1, 2, 3, 3},
			wantMaxSizes: []int{1, 2, 2, 2, 2},
		},
		{
			name: "Single island",
			args: args{
				m:         1,
				n:         1,
				positions: [][]int{{0, 0}},
			},
			want:         []int{1},
			wantMaxSizes: []int{1},
		},
		{
			name: "Merging two separated islands",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {0, 2}, {0, 1}},
			},
			// {0,0} -> c=1, max=1
			// {0,2} -> c=2, max=1 (two islands of size 1)
			// {0,1} -> connects {0,0} and {0,2} -> c=1, max=3
			want:         []int{1, 2, 1},
			wantMaxSizes: []int{1, 1, 3},
		},
		{
			name: "Large Disconnected",
			args: args{
				m:         10,
				n:         10,
				positions: [][]int{{0, 0}, {5, 5}, {9, 9}},
			},
			want:         []int{1, 2, 3},
			wantMaxSizes: []int{1, 1, 1},
		},
		{
			name: "L-Shape formation",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}},
			},
			// {0,0} -> 1, max=1
			// {1,0} -> union({0,0}), size=2 -> c=1, max=2
			// {2,0} -> union({1,0}), size=3 -> c=1, max=3
			// {2,1} -> union({2,0}), size=4 -> c=1, max=4
			want:         []int{1, 1, 1, 1},
			wantMaxSizes: []int{1, 2, 3, 4},
		},
		{
			name: "Merge 4 islands (Cross)",
			args: args{
				m:         3,
				n:         3,
				positions: [][]int{{1, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 1}},
			},
			// {1,0}: c=1, m=1
			// {0,1}: c=2, m=1
			// {1,2}: c=3, m=1
			// {2,1}: c=4, m=1
			// {1,1}: connects all 4. c=1, m=5
			want:         []int{1, 2, 3, 4, 1},
			wantMaxSizes: []int{1, 1, 1, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotMaxSizes := numIslands2(tt.args.m, tt.args.n, tt.args.positions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotMaxSizes, tt.wantMaxSizes) {
				t.Errorf("numIslands2() gotMaxSizes = %v, want %v", gotMaxSizes, tt.wantMaxSizes)
			}
		})
	}
}
