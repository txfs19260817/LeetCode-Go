package leetcode

import "testing"

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "babgbag", t = "bag"`,
			args: args{"babgbag", "bag"},
			want: 5,
		},
		{
			name: `s = "acbb", t = "ab"`,
			args: args{"acbb", "ab"},
			want: 2,
		},
		{
			name: `s = "rabbbit", t = "rabbit"`,
			args: args{"rabbbit", "rabbit"},
			want: 3,
		}, {
			name: `s = "ccc", t = "c"`,
			args: args{"ccc", "c"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
