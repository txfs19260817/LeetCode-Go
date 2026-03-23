package leetcode

import "testing"

func Test_canTransform(t *testing.T) {
	type args struct {
		start  string
		result string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				start:  "RXXLRXRXL",
				result: "XRLXXRRLX",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				start:  "X",
				result: "L",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTransform(tt.args.start, tt.args.result); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
