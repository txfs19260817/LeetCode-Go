package _160_Find_Words_That_Can_Be_Formed_by_Characters

import "testing"

func Test_countCharacters(t *testing.T) {
	type args struct {
		words []string
		chars string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `words = ["cat","bt","hat","tree"], chars = "atach"`,
			args: args{[]string{"cat", "bt", "hat", "tree"}, "atach"},
			want: 6,
		},
		{
			name: `words = ["hello","world","leetcode"], chars = "welldonehoneyr"`,
			args: args{[]string{"hello", "world", "leetcode"}, "welldonehoneyr"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCharacters(tt.args.words, tt.args.chars); got != tt.want {
				t.Errorf("countCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
