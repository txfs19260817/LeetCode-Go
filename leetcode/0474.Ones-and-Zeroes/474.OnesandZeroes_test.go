package _474_Ones_and_Zeroes

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `strs = ["10","0001","111001","1","0"], m = 5, n = 3`,
			args: args{[]string{"10", "0001", "111001", "1", "0"}, 5, 3},
			want: 4,
		},
		{
			name: `strs = ["10","0","1"], m = 1, n = 1`,
			args: args{[]string{"10", "0", "1"}, 1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
