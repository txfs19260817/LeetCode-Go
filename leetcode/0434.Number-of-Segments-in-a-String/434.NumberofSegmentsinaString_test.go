package leetcode

import "testing"

func Test_countSegments(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `"Hello, my name is John"`,
			args: args{"Hello, my name is John"},
			want: 5,
		},
		{
			name: `"Hello,"`,
			args: args{"Hello,"},
			want: 1,
		},
		{
			name: "",
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
