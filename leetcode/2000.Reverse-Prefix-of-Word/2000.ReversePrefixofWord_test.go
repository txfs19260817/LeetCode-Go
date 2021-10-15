package _000_Reverse_Prefix_of_Word

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
			name: `word = "abcdefd", ch = "z"`,
			args: args{"abcdefd", 'z'},
			want: "abcdefd",
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
