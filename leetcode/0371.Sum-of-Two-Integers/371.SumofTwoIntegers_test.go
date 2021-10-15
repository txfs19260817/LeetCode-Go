package _371_Sum_of_Two_Integers

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a = 2, b = 3",
			args: args{2, 3},
			want: 5,
		},
		{
			name: "a = 20, b = 30",
			args: args{20, 30},
			want: 50,
		},
		{
			name: "a = -12, b = -8",
			args: args{-12, -8},
			want: -20,
		},
		{
			name: "a = -12, b = 8",
			args: args{-12, 8},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
