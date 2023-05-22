package leetcode

import "testing"

func Test_getSmallestString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "n = 3, k = 27",
			args: args{3, 27},
			want: "aay",
		},
		{
			name: "n = 24, k = 552",
			args: args{24, 552},
			want: "aadzzzzzzzzzzzzzzzzzzzzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
