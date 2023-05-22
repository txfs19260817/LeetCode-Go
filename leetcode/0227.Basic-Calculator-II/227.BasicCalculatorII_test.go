package leetcode

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3+2*2",
			args: args{"3+2*2"},
			want: 7,
		},
		{
			name: " 3/2 ",
			args: args{" 3/2 "},
			want: 1,
		},
		{
			name: " 3+5 / 2 ",
			args: args{" 3+5 / 2 "},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
