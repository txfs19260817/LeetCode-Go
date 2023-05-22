package leetcode

import "testing"

func Test_minimizeResult(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "247+38",
			args: args{"247+38"},
			want: "2(47+38)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeResult(tt.args.expression); got != tt.want {
				t.Errorf("minimizeResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
