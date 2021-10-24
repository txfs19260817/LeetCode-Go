package _906_Number_of_Valid_Words_in_a_Sentence

import "testing"

func Test_countValidWords(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `sentence = "cat and  dog"`,
			args: args{"cat and  dog"},
			want: 3,
		},
		{
			name: `sentence = "!this  1-s b8d!"`,
			args: args{"!this  1-s b8d!"},
			want: 0,
		},
		{
			name: `sentence = "alice and  bob are playing stone-game10"`,
			args: args{"alice and  bob are playing stone-game10"},
			want: 5,
		},
		{
			name: `sentence = "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."`,
			args: args{"he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidWords(tt.args.sentence); got != tt.want {
				t.Errorf("countValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
