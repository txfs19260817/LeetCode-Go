package _234_Find_Resultant_Array_After_Removing_Anagrams

import (
	"reflect"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `"abba","baba","bbaa","cd","cd"`,
			args: args{[]string{"abba", "baba", "bbaa", "cd", "cd"}},
			want: []string{"abba", "cd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAnagrams(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
