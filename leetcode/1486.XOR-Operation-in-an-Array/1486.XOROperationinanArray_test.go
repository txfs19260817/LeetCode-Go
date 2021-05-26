package _486_XOR_Operation_in_an_Array

import "testing"

func Test_xorOperation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 5, start = 0",
			args: args{5, 0},
			want: 8,
		},
		{
			name: "n = 4, start = 3",
			args: args{4, 3},
			want: 8,
		},
		{
			name: "n = 1, start = 7",
			args: args{1, 7},
			want: 7,
		},
		{
			name: "n = 10, start = 5",
			args: args{10, 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
