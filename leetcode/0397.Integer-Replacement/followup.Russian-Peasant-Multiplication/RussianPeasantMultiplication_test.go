package leetcode

import "testing"

func Test_russianPeasantMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "example: 13 * 12",
			a:    13,
			b:    12,
			want: 156,
		},
		{
			name: "zero times positive",
			a:    0,
			b:    7,
			want: 0,
		},
		{
			name: "positive times zero",
			a:    37,
			b:    0,
			want: 0,
		},
		{
			name: "negative times positive",
			a:    -13,
			b:    12,
			want: -156,
		},
		{
			name: "positive times negative",
			a:    13,
			b:    -12,
			want: -156,
		},
		{
			name: "negative times negative",
			a:    -13,
			b:    -12,
			want: 156,
		},
		{
			name: "odd times odd",
			a:    15,
			b:    15,
			want: 225,
		},
		{
			name: "power of two factor",
			a:    64,
			b:    3,
			want: 192,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := russianPeasantMultiply(tt.a, tt.b); got != tt.want {
				t.Errorf("russianPeasantMultiply() = %v, want %v", got, tt.want)
			}
			if got := russianPeasantMultiplyIterative(tt.a, tt.b); got != tt.want {
				t.Errorf("russianPeasantMultiplyIterative() = %v, want %v", got, tt.want)
			}
			if got := russianPeasantMultiplyRecursive(tt.a, tt.b); got != tt.want {
				t.Errorf("russianPeasantMultiplyRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
