package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "34.00515, -3",
			args: args{34.00515, -3},
			want: 0.00003,
		},
		{
			name: "3, 3",
			args: args{3, 3},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got > tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
			if got := myPow1(tt.args.x, tt.args.n); got > tt.want {
				t.Errorf("myPow1() = %v, want %v", got, tt.want)
			}
		})
	}
}
