package k

import (
	"reflect"
	"testing"
)

func Test_findWordLocation(t *testing.T) {
	grid1 := [][]byte{
		{'c', 'c', 'x', 't', 'i', 'b'},
		{'c', 'c', 'a', 't', 'n', 'i'},
		{'a', 'c', 'n', 'n', 't', 't'},
		{'t', 'c', 's', 'i', 'p', 't'},
		{'a', 'o', 'o', 'o', 'a', 'a'},
		{'o', 'a', 'a', 'a', 'o', 'o'},
		{'k', 'a', 'i', 'c', 'k', 'i'},
	}
	type args struct {
		grid [][]byte
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "catnip",
			args:    args{grid1, "catnip"},
			wantAns: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}, {3, 4}},
		},
		{
			name:    "s",
			args:    args{grid1, "s"},
			wantAns: [][]int{{3, 2}},
		},
		{
			name:    "bit",
			args:    args{grid1, "bit"},
			wantAns: [][]int{{0, 5}, {1, 5}, {2, 5}},
		},
		{
			name:    "aoi",
			args:    args{grid1, "aoi"},
			wantAns: [][]int{{4, 5}, {5, 5}, {6, 5}},
		},
		{
			name:    "ki",
			args:    args{grid1, "ki"},
			wantAns: [][]int{{6, 4}, {6, 5}},
		},
		{
			name:    "aaa",
			args:    args{grid1, "aaa"},
			wantAns: [][]int{{5, 1}, {5, 2}, {5, 3}},
		},
		{
			name:    "ooo",
			args:    args{grid1, "ooo"},
			wantAns: [][]int{{4, 1}, {4, 2}, {4, 3}},
		},
		{
			name:    "a",
			args:    args{[][]byte{{'a'}}, "a"},
			wantAns: [][]int{{0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findWordLocation(tt.args.grid, tt.args.word); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findWordLocation() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
