package _342_Power_of_Four

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "16",
			args: args{16},
			want: true,
		},
		{
			name: "5",
			args: args{5},
			want: false,
		},
		{
			name: "1",
			args: args{1},
			want: true,
		},
		{
			name: "-1",
			args: args{-1},
			want: false,
		},
		{
			name: "0",
			args: args{0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
			if got := isPowerOfFour1(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour1() = %v, want %v", got, tt.want)
			}
		})
	}
}
