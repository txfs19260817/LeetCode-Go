package _030_Substring_with_Concatenation_of_All_Words

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "barfoothefoobarman",
			args: args{"barfoothefoobarman", []string{"foo", "bar"}},
			want: []int{0, 9},
		},
		{
			name: "abc",
			args: args{"abc", []string{"a", "b", "c"}},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	}
}

func Benchmark_findSubstring1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findSubstring1("barfoothefoobarman", []string{"foo", "bar"})
	}
}
