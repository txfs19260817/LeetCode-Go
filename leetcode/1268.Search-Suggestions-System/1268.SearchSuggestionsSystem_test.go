package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: `products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"`,
			args: args{[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"},
			want: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			name: "havana",
			args: args{[]string{"havana"}, "havana"},
			want: [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			name: `"code","codephone","coddle","coddles","codes"`,
			args: args{[]string{"code", "codephone", "coddle", "coddles", "codes"}, "coddle"},
			want: [][]string{
				{"coddle", "coddles", "code"},
				{"coddle", "coddles", "code"},
				{"coddle", "coddles", "code"},
				{"coddle", "coddles"},
				{"coddle", "coddles"},
				{"coddle", "coddles"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, suggestedProducts(tt.args.products, tt.args.searchWord))
			assert.Equal(t, tt.want, suggestedProducts2(tt.args.products, tt.args.searchWord))
		})
	}
}

func Benchmark_suggestedProducts(b *testing.B) {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	for b.Loop() {
		_ = suggestedProducts(products, "mouse")
	}
}

func Benchmark_suggestedProducts2(b *testing.B) {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	for b.Loop() {
		benchmarkProducts := append([]string(nil), products...)
		_ = suggestedProducts2(benchmarkProducts, "mouse")
	}
}
