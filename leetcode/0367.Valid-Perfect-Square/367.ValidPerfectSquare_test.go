package leetcode

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "14",
			args: args{14},
			want: false,
		},
		{
			name: "16",
			args: args{16},
			want: true,
		},
		{
			name: "1",
			args: args{1},
			want: true,
		},
		{
			name: "99980001",
			args: args{99980001},
			want: true,
		},
		{
			name: "99980000",
			args: args{99980000},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
			if got := isPerfectSquare2(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare2() = %v, want %v", got, tt.want)
			}
		})
	}
}
