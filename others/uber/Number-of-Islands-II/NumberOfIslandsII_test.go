package uber

import (
	"reflect"
	"testing"
)

func TestNumberOfIslandsII(t *testing.T) {
	tests := []struct {
		name      string
		m         int
		n         int
		positions [][]int
		want      []int
	}{
		{
			name:      "example-1",
			m:         3,
			n:         3,
			positions: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}},
			want:      []int{1, 1, 2, 3},
		},
		{
			name:      "example-2",
			m:         1,
			n:         1,
			positions: [][]int{{0, 0}},
			want:      []int{1},
		},
		{
			name:      "duplicate-positions",
			m:         2,
			n:         2,
			positions: [][]int{{0, 0}, {0, 0}, {1, 1}},
			want:      []int{1, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfIslandsII(tt.m, tt.n, tt.positions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("NumberOfIslandsII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOfIslandsIIMaxSize(t *testing.T) {
	m := 3
	n := 3
	positions := [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 1}}

	wantCounts := []int{1, 1, 2, 3, 1}
	wantMax := []int{1, 2, 2, 2, 5}

	gotCounts, gotMax := NumberOfIslandsIIMaxSize(m, n, positions)
	if !reflect.DeepEqual(gotCounts, wantCounts) {
		t.Fatalf("NumberOfIslandsIIMaxSize counts = %v, want %v", gotCounts, wantCounts)
	}
	if !reflect.DeepEqual(gotMax, wantMax) {
		t.Fatalf("NumberOfIslandsIIMaxSize max sizes = %v, want %v", gotMax, wantMax)
	}
}
