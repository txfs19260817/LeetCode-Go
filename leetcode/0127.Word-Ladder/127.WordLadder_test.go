package _127_Word_Ladder

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `"hot","dot","dog","lot","log","cog"`,
			args: args{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			want: 5,
		},
		{
			name: `"hot","dot","dog","lot","log"`,
			args: args{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
