package leetcode

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4`,
			args: args{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4},
			want: []string{"the", "is", "sunny", "day"},
		},
		{
			name: `["i", "love", "leetcode", "i", "love", "coding"], k = 2`,
			args: args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2},
			want: []string{"i", "love"},
		},
		{
			name: `["i", "love", "leetcode", "i", "love", "coding"], k = 1`,
			args: args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 1},
			want: []string{"i"},
		},
		{
			name: `["i", "love", "leetcode", "i", "love", "coding"], k = 3`,
			args: args{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 3},
			want: []string{"i", "love", "coding"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
