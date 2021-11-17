package _318_Maximum_Product_of_Word_Lengths

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `words = ["abcw","baz","foo","bar","xtfn","abcdef"]`,
			args: args{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			want: 16,
		},
		{
			name: `words = ["a","ab","abc","d","cd","bcd","abcd"]`,
			args: args{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			want: 4,
		},
		{
			name: `words = ["a","aa","aaa","aaaa"]`,
			args: args{[]string{"a", "aa", "aaa", "aaaa"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
			if got := maxProduct2(tt.args.words); got != tt.want {
				t.Errorf("maxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
	}
}

func Benchmark_maxProduct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProduct2([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
	}
}
