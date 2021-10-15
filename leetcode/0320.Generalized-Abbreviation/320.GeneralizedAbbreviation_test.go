package _320_Generalized_Abbreviation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateAbbreviations(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "word",
			args: args{"word"},
			want: []string{"4", "3d", "2r1", "2rd", "1o2", "1o1d", "1or1", "1ord", "w3", "w2d", "w1r1", "w1rd", "wo2", "wo1d", "wor1", "word"},
		},
		{
			name: "a",
			args: args{"a"},
			want: []string{"1", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, generateAbbreviations(tt.args.word))
		})
	}
}
