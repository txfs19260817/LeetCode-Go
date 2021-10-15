package _318_Minimum_Flips_to_Make_a_OR_b_Equal_to_c

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a = 2, b = 6, c = 5",
			args: args{2, 6, 5},
			want: 3,
		},
		{
			name: "a = 4, b = 2, c = 7",
			args: args{4, 2, 7},
			want: 1,
		},
		{
			name: "a = 1, b = 2, c = 3",
			args: args{1, 2, 3},
			want: 0,
		},
		{
			name: "a = 8, b = 3, c = 5",
			args: args{8, 3, 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
