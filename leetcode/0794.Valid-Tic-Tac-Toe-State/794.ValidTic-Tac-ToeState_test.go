package leetcode

import "testing"

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{board: []string{"O  ", "   ", "   "}},
			want: false,
		},
		{
			name: "Example 2",
			args: args{board: []string{"XOX", " X ", "   "}},
			want: false,
		},
		{
			name: "Example 3",
			args: args{board: []string{"XOX", "O O", "XOX"}},
			want: true,
		},
		{
			name: "XXX and OOO",
			args: args{board: []string{"XXX", "   ", "OOO"}},
			want: false,
		},
		{
			name: "XXX with OOX twice",
			args: args{board: []string{"XXX", "OOX", "OOX"}},
			want: true,
		},
		{
			name: "XXX with XOO and OO",
			args: args{board: []string{"XXX", "XOO", "OO "}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTicTacToe(tt.args.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}
