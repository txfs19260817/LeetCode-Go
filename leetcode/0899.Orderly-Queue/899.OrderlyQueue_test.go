package leetcode

import "testing"

func Test_orderlyQueue(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `s = "cba", k = 1`,
			args: args{"cba", 1},
			want: "acb",
		},
		{
			name: `s = "baaca", k = 3`,
			args: args{"baaca", 3},
			want: "aaabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderlyQueue(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("orderlyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
