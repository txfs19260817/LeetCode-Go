package _704_Determine_if_String_Halves_Are_Alike

import "testing"

func Test_halvesAreAlike(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s = "book"`,
			args: args{"book"},
			want: true,
		},
		{
			name: `s = "AbCdEfGh"`,
			args: args{"AbCdEfGh"},
			want: true,
		},
		{
			name: `s = "textbook"`,
			args: args{"textbook"},
			want: false,
		},
		{
			name: `s = "MerryChristmas"`,
			args: args{"MerryChristmas"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlike(tt.args.s); got != tt.want {
				t.Errorf("halvesAreAlike() = %v, want %v", got, tt.want)
			}
		})
	}
}
