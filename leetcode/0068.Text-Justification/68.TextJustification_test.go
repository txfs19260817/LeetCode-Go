package _068_Text_Justification

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16`,
			args: args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
			want: []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name: `words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16`,
			args: args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
			want: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		}, {
			name: `words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20`,
			args: args{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20},
			want: []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fullJustify(tt.args.words, tt.args.maxWidth)
			assert.Len(t, got, len(tt.want))
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], got[i])
			}
		})
	}
}
