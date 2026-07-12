package leetcode

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	type args struct {
		board [][]byte
		click []int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "reveals blank area recursively",
			args: args{
				board: boardFromRows(
					"EEEEE",
					"EEMEE",
					"EEEEE",
					"EEEEE",
				),
				click: []int{3, 0},
			},
			want: boardFromRows(
				"B1E1B",
				"B1M1B",
				"B111B",
				"BBBBB",
			),
		},
		{
			name: "reveals clicked mine",
			args: args{
				board: boardFromRows(
					"B1E1B",
					"B1M1B",
					"B111B",
					"BBBBB",
				),
				click: []int{1, 2},
			},
			want: boardFromRows(
				"B1E1B",
				"B1X1B",
				"B111B",
				"BBBBB",
			),
		},
		{
			name: "reveals adjacent mine count",
			args: args{
				board: boardFromRows(
					"EM",
				),
				click: []int{0, 0},
			},
			want: boardFromRows(
				"1M",
			),
		},
		{
			name: "reveals all blank cells when no mines exist",
			args: args{
				board: boardFromRows(
					"EE",
					"EE",
				),
				click: []int{0, 0},
			},
			want: boardFromRows(
				"BB",
				"BB",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard() = %v, want %v", got, tt.want)
			}
			if got := updateBoard2(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_updateBoard(b *testing.B) {
	benchmarkUpdateBoard(b, updateBoard)
}

func Benchmark_updateBoard2(b *testing.B) {
	benchmarkUpdateBoard(b, updateBoard2)
}

func benchmarkUpdateBoard(b *testing.B, update func([][]byte, []int) [][]byte) {
	board := boardFromRows(
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
		"EEEEEEEEEEEEEEEE",
	)
	click := []int{0, 0}

	b.ReportAllocs()
	b.StopTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boardCopy := cloneBoard(board)
		b.StartTimer()
		_ = update(boardCopy, click)
		b.StopTimer()
	}
}

func boardFromRows(rows ...string) [][]byte {
	board := make([][]byte, len(rows))
	for i, row := range rows {
		board[i] = []byte(row)
	}
	return board
}

func cloneBoard(board [][]byte) [][]byte {
	clone := make([][]byte, len(board))
	for i := range board {
		clone[i] = append([]byte(nil), board[i]...)
	}
	return clone
}
