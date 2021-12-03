package k

import (
	"reflect"
	"testing"
)

func Test_one(t *testing.T) {
	type args struct {
		mat [][]byte
	}
	tests := []struct {
		name       string
		args       args
		wantStart  []int
		wantEnd    []int
		wantHeight int
		wantWidth  int
	}{
		{
			name:       "1",
			args:       args{[][]byte{{'1', '1', '1', '1', '1'}, {'1', '1', '0', '0', '1'}, {'1', '1', '0', '0', '1'}, {'1', '1', '1', '1', '1'}}},
			wantStart:  []int{1, 2},
			wantEnd:    []int{2, 3},
			wantHeight: 2,
			wantWidth:  2,
		},
		{
			name:       "2",
			args:       args{[][]byte{{'1', '1', '1', '1', '1'}, {'1', '1', '0', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}}},
			wantStart:  []int{1, 2},
			wantEnd:    []int{1, 2},
			wantHeight: 1,
			wantWidth:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, gotHeight, gotWidth := one(tt.args.mat)
			if !reflect.DeepEqual(gotStart, tt.wantStart) {
				t.Errorf("one() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if !reflect.DeepEqual(gotEnd, tt.wantEnd) {
				t.Errorf("one() gotEnd = %v, want %v", gotEnd, tt.wantEnd)
			}
			if gotHeight != tt.wantHeight {
				t.Errorf("one() gotHeight = %v, want %v", gotHeight, tt.wantHeight)
			}
			if gotWidth != tt.wantWidth {
				t.Errorf("one() gotWidth = %v, want %v", gotWidth, tt.wantWidth)
			}
		})
	}
}

func Test_multiple(t *testing.T) {
	type args struct {
		mat [][]byte
	}
	tests := []struct {
		name       string
		args       args
		wantStarts [][]int
		wantEnds   [][]int
	}{
		{
			name:       "1",
			args:       args{[][]byte{{'0', '1', '1', '1', '1'}, {'1', '1', '0', '0', '1'}, {'0', '1', '0', '0', '1'}, {'0', '1', '1', '1', '1'}, {'1', '0', '1', '1', '1'}}},
			wantStarts: [][]int{{0, 0}, {1, 2}, {2, 0}, {4, 1}},
			wantEnds:   [][]int{{0, 0}, {2, 3}, {3, 0}, {4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStarts, gotEnds := multiple(tt.args.mat)
			if !reflect.DeepEqual(gotStarts, tt.wantStarts) {
				t.Errorf("multiple() gotStarts = %v, want %v", gotStarts, tt.wantStarts)
			}
			if !reflect.DeepEqual(gotEnds, tt.wantEnds) {
				t.Errorf("multiple() gotEnds = %v, want %v", gotEnds, tt.wantEnds)
			}
		})
	}
}
