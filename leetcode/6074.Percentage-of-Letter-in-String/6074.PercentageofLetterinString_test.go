package _074_Percentage_of_Letter_in_String

import "testing"

func Test_percentageLetter(t *testing.T) {
	type args struct {
		s      string
		letter byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "foobar", letter = "o"`,
			args: args{"foobar", 'o'},
			want: 33,
		},
		{
			name: `s = "foobar", letter = "z"`,
			args: args{"foobar", 'z'},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentageLetter(tt.args.s, tt.args.letter); got != tt.want {
				t.Errorf("percentageLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
