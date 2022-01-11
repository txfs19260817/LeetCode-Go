package _408_Valid_Word_Abbreviation

import "testing"

func Test_validWordAbbreviation(t *testing.T) {
	type args struct {
		word string
		abbr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `word = "internationalization", abbr = "i12iz4n"`,
			args: args{"internationalization", "i12iz4n"},
			want: true,
		},
		{
			name: `word = "apple", abbr = "a2e"`,
			args: args{"apple", "a2e"},
			want: false,
		},
		{
			name: `word = "a", abbr = "01"`,
			args: args{"a", "01"},
			want: false,
		},
		{
			name: `word = "hi", abbr = "1"`,
			args: args{"hi", "1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validWordAbbreviation(tt.args.word, tt.args.abbr); got != tt.want {
				t.Errorf("validWordAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
