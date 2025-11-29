package k

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPassableLanes(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name     string
		args     args
		wantRows []int
		wantCols []int
	}{
		{
			name: "board1",
			args: args{
				board: [][]byte{
					{'+', '+', '+', '0', '+', '0', '0'},
					{'0', '0', '+', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '+', '0', '0'},
					{'+', '+', '+', '0', '0', '+', '0'},
					{'0', '0', '0', '0', '0', '0', '0'},
				},
			},
			wantRows: []int{4},
			wantCols: []int{3, 6},
		},
		{
			name: "board2",
			args: args{
				board: [][]byte{
					{'+', '+', '+', '0', '+', '0', '0'},
					{'0', '0', '0', '0', '0', '+', '0'},
					{'0', '0', '+', '0', '0', '0', '0'},
					{'0', '0', '0', '0', '+', '0', '0'},
					{'+', '+', '+', '0', '0', '0', '+'},
				},
			},
			wantRows: nil, // or []int{} depending on implementation, usually nil if empty
			wantCols: []int{3},
		},
		{
			name: "board3",
			args: args{
				board: [][]byte{
					{'+', '+', '+', '0', '+', '0', '0'},
					{'0', '0', '0', '0', '0', '0', '0'},
					{'0', '0', '+', '+', '0', '+', '0'},
					{'0', '0', '0', '0', '+', '0', '0'},
					{'+', '+', '+', '0', '0', '0', '+'},
				},
			},
			wantRows: []int{1},
			wantCols: nil,
		},
		{
			name: "board4",
			args: args{
				board: [][]byte{{'+'}},
			},
			wantRows: nil,
			wantCols: nil,
		},
		{
			name: "board5",
			args: args{
				board: [][]byte{{'0'}},
			},
			wantRows: []int{0},
			wantCols: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, gotCols := findPassableLanes(tt.args.board)
			assert.ElementsMatch(t, gotRows, tt.wantRows)
			assert.ElementsMatch(t, gotCols, tt.wantCols)
		})
	}
}

func Test_findExit(t *testing.T) {
	type args struct {
		board [][]byte
		start [2]int
	}

	// Helper to clone board because findExit might not modify it but test cases share definitions in comments
	// Actually our implementation does not modify board (uses visited array), so it's safe.

	board1 := [][]byte{
		{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '+', '+'},
		{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
		{'+', '+', '0', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
		{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
	}

	board2 := [][]byte{
		{'+', '+', '+', '+', '+', '+', '+'},
		{'0', '0', '0', '0', '+', '0', '+'},
		{'+', '0', '+', '0', '+', '0', '0'},
		{'+', '0', '0', '0', '+', '+', '+'},
		{'+', '+', '+', '+', '+', '+', '+'},
	}

	board3 := [][]byte{
		{'+', '0', '+', '0', '+'},
		{'0', '0', '+', '0', '0'},
		{'+', '0', '+', '0', '+'},
		{'0', '0', '+', '0', '0'},
		{'+', '0', '+', '0', '+'},
	}

	board4 := [][]byte{
		{'+', '0', '+', '0', '+'},
		{'0', '0', '0', '0', '0'},
		{'+', '+', '+', '+', '+'},
		{'0', '0', '0', '0', '0'},
		{'+', '0', '+', '0', '+'},
	}

	board5 := [][]byte{
		{'+', '0', '0', '0', '+'},
		{'+', '0', '+', '0', '+'},
		{'+', '0', '0', '0', '+'},
		{'+', '0', '+', '0', '+'},
	}

	tests := []struct {
		name    string
		args    args
		wantEnd [2]int
	}{
		{"board1_start1", args{board1, [2]int{2, 0}}, [2]int{5, 2}},
		{"board1_start2", args{board1, [2]int{0, 7}}, [2]int{0, 8}},
		{"board1_start3", args{board1, [2]int{5, 2}}, [2]int{2, 0}}, // 2,0 or 5,5 -> 2,0 preferred
		{"board1_start4", args{board1, [2]int{5, 5}}, [2]int{5, 7}},
		{"board2_start1", args{board2, [2]int{1, 0}}, [2]int{-1, -1}},
		{"board2_start2", args{board2, [2]int{2, 6}}, [2]int{-1, -1}},
		{"board3_start1", args{board3, [2]int{0, 1}}, [2]int{1, 0}},
		{"board3_start2", args{board3, [2]int{4, 1}}, [2]int{3, 0}},
		{"board3_start3", args{board3, [2]int{0, 3}}, [2]int{1, 4}},
		{"board3_start4", args{board3, [2]int{4, 3}}, [2]int{3, 4}},
		{"board4_start1", args{board4, [2]int{1, 0}}, [2]int{0, 1}},
		{"board4_start2", args{board4, [2]int{1, 4}}, [2]int{0, 3}},
		{"board4_start3", args{board4, [2]int{3, 0}}, [2]int{4, 1}},
		{"board4_start4", args{board4, [2]int{3, 4}}, [2]int{4, 3}},
		{"board5_start1", args{board5, [2]int{0, 1}}, [2]int{0, 2}},
		{"board5_start2", args{board5, [2]int{3, 1}}, [2]int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create deep copy of board
			boardCopy := make([][]byte, len(tt.args.board))
			for i := range tt.args.board {
				boardCopy[i] = make([]byte, len(tt.args.board[i]))
				copy(boardCopy[i], tt.args.board[i])
			}

			got := findExit(boardCopy, tt.args.start)
			if !reflect.DeepEqual(got, tt.wantEnd) {
				t.Errorf("findExit() = %v, want %v", got, tt.wantEnd)
			}
		})
	}
}

func Test_getNestEntranceCount(t *testing.T) {
	type args struct {
		board [][]byte
	}

	board1 := [][]byte{
		{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '+', '+'},
		{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
		{'+', '+', '0', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
		{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
	}
	board2 := [][]byte{
		{'+', '+', '+', '+', '+', '+'},
		{'0', '0', '0', '+', '0', '+'},
		{'+', '0', '+', '0', '0', '0'},
		{'+', '+', '+', '+', '+', '+'},
	}
	board3 := [][]byte{
		{'+', '0', '+', '+', '+', '0', '+', '0', '0'},
		{'+', '0', '+', '0', '0', '0', '0', '+', '+'},
		{'0', '0', '0', '0', '0', '+', '+', '0', '+'},
		{'+', '+', '+', '+', '+', '+', '+', '0', '0'},
		{'+', '+', '0', '0', '0', '0', '0', '0', '+'},
		{'+', '+', '0', '+', '+', '0', '+', '0', '+'},
	}
	board4 := [][]byte{
		{'+', '+', '+'},
		{'+', '0', '+'},
		{'+', '+', '+'},
	}
	board5 := [][]byte{{'+'}}
	board6 := [][]byte{
		{'0', '0'},
		{'0', '0'},
	}

	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{"board1", args{board1}, []int{2, 5}},
		{"board2", args{board2}, []int{1, 1}},
		{"board3", args{board3}, []int{2, 4, 3}},
		{"board4", args{board4}, []int{0}},
		{"board5", args{board5}, nil},
		{"board6", args{board6}, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNestEntranceCount(tt.args.board)
			assert.ElementsMatch(t, tt.wantAns, got, "getNestEntranceCount(%v)", tt.args.board)
		})
	}
}
