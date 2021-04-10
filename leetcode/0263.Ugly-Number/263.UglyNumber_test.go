package _263_Ugly_Number

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n = 6",
			args: args{6},
			want: true,
		},
		{
			name: "n = 1",
			args: args{1},
			want: true,
		},
		{
			name: "n = 8",
			args: args{8},
			want: true,
		},
		{
			name: "n = 14",
			args: args{14},
			want: false,
		},
		{
			name: "n = 0",
			args: args{0},
			want: false,
		},
		{
			name: "n = -1",
			args: args{-1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
