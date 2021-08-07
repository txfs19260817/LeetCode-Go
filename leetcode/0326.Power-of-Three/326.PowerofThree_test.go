package _326_Power_of_Three

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "27",
			args: args{27},
			want: true,
		},
		{
			name: "0",
			args: args{0},
			want: false,
		},
		{
			name: "1",
			args: args{1},
			want: true,
		},
		{
			name: "9",
			args: args{9},
			want: true,
		},
		{
			name: "2",
			args: args{2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
