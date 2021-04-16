package _953_Verifying_an_Alien_Dictionary

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"`,
			args: args{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: `words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"`,
			args: args{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"},
			want: false,
		},
		{
			name: `words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"`,
			args: args{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
