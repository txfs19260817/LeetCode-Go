package _614_Maximum_Nesting_Depth_of_the_Parentheses

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "(1+(2*3)+((8)/4))+1",
			args: args{"(1+(2*3)+((8)/4))+1"},
			want: 3,
		},
		{
			name: "",
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.s); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
