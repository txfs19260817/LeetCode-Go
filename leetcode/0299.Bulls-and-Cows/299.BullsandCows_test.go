package leetcode

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `secret = "1807", guess = "7810"`,
			args: args{"1807", "7810"},
			want: "1A3B",
		},
		{
			name: `secret = "1123", guess = "0111"`,
			args: args{"1123", "0111"},
			want: "1A1B",
		},
		{
			name: `secret = "1", guess = "1"`,
			args: args{"1", "1"},
			want: "1A0B",
		},
		{
			name: `secret = "1", guess = "0"`,
			args: args{"1", "0"},
			want: "0A0B",
		},
		{
			name: `secret = "00112233445566778899", guess = "16872590340158679432"`,
			args: args{"00112233445566778899", "16872590340158679432"},
			want: "3A17B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
