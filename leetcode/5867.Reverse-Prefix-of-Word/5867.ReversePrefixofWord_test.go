package _867_Reverse_Prefix_of_Word

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `word = "abcdefd", ch = "d"`,
			args: args{"abcdefd", 'd'},
			want: "dcbaefd",
		},
		{
			name: `word = "xyxzxe", ch = "z"`,
			args: args{"xyxzxe", 'z'},
			want: "zxyxxe",
		},
		{
			name: `word = "abcd", ch = "z"`,
			args: args{"abcd", 'z'},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
