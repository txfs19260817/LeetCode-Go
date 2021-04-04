package _036_Valid_Sudoku

import "testing"

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{[][]byte{
				[]byte(".87654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			want: true,
		},
		{
			name: "2",
			args: args{[][]byte{
				[]byte(".87654321"),
				[]byte("2.3......"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			want: false,
		},
		{
			name: "3",
			args: args{[][]byte{
				[]byte(".88654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			want: false,
		},
		{
			name: "4",
			args: args{[][]byte{
				[]byte(".87654321"),
				[]byte("2....4..."),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
