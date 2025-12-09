package leetcode

import (
	"reflect"
	"testing"
)

func Test_wordsAbbreviation(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"},
			},
			want: []string{"l2e", "god", "internal", "me", "i6t", "interval", "inte4n", "f2e", "intr4n"},
		},
		{
			name: "Example 2",
			args: args{
				words: []string{"aa", "aaa"},
			},
			want: []string{"aa", "aaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordsAbbreviation(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordsAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
