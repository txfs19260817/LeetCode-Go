package leetcode

import "testing"

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: ".L.R...LR..L..",
			args: args{".L.R...LR..L.."},
			want: "LL.RR.LLRRLL..",
		},
		{
			name: "RR.L",
			args: args{"RR.L"},
			want: "RR.L",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushDominoes(tt.args.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
