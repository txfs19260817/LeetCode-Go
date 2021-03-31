package _028_Implement_strStr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `haystack = "hello", needle = "ll"`,
			args: args{"hello", "ll"},
			want: 2,
		},
		{
			name: `haystack = "aaaaa", needle = "bba"`,
			args: args{"aaaaa", "bba"},
			want: -1,
		},
		{
			name: `haystack = "", needle = ""`,
			args: args{"", ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
