package k

import (
	"reflect"
	"testing"
)

var board = [][]int{
	{0, 0, 0, -1, -1},
	{0, 0, -1, 0, 0},
	{0, -1, 0, -1, 0},
	{0, 0, -1, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0}}

var board2 = [][]int{{0, 0, 0, 0, -1}, {0, -1, -1, 0, 0}, {0, 0, 0, 0, 0}, {-1, -1, 0, 0, 0}, {0, -1, 0, 0, 0}, {0, -1, 0, 0, 0}}
var board25 = [][]int{{0, 0, 0, 0, -1}, {0, -1, -1, 0, 0}, {0, 0, 0, 0, 0}, {-1, -1, 0, 0, 0}, {0, -1, 0, 0, 0}, {0, 0, 0, 0, 0}}
var board3 = [][]int{{1, 0, 0, 0, 0}, {0, -1, -1, 0, 0}, {0, -1, 0, 1, 0}, {-1, 0, 0, 0, 0}, {0, 1, -1, 0, 0}, {0, 0, 0, 0, 0}}

func Test_findLegalMoves(t *testing.T) {
	type args struct {
		mat [][]int
		i   int
		j   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{board, 2, 2},
			want: nil,
		},
		{
			name: "2",
			args: args{board, 0, 0},
			want: [][]int{{0, 1}, {1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLegalMoves(tt.args.mat, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLegalMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAllPositionsPassed(t *testing.T) {
	type args struct {
		board [][]int
		endI  int
		endJ  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{board, 0, 0},
			want: false,
		},
		{
			name: "2",
			args: args{board2, 0, 0},
			want: false,
		},
		{
			name: "2.5",
			args: args{board25, 0, 0},
			want: true,
		},
		{
			name: "3",
			args: args{board3, 1, 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllPositionsPassed(tt.args.board, tt.args.endI, tt.args.endJ); got != tt.want {
				t.Errorf("isAllPositionsPassed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAllTreasures(t *testing.T) {
	type args struct {
		board [][]int
		start []int
		end   []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name:    "1",
			args:    args{board3, []int{5, 0}, []int{0, 4}},
			wantAns: nil,
		},
		{
			name:    "2",
			args:    args{board3, []int{5, 2}, []int{2, 0}},
			wantAns: [][]int{{5, 2}, {5, 1}, {4, 1}, {3, 1}, {3, 2}, {3, 3}, {2, 3}, {1, 3}, {0, 3}, {0, 2}, {0, 1}, {0, 0}, {1, 0}, {2, 0}},
		},
		{
			name:    "3",
			args:    args{board3, []int{0, 0}, []int{4, 1}},
			wantAns: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 3}, {3, 3}, {3, 2}, {3, 1}, {4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAllTreasures(tt.args.board, tt.args.start, tt.args.end); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAllTreasures() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
