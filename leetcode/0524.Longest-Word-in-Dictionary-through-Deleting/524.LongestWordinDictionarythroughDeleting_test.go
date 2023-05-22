package leetcode

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s string
		d []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]`,
			args: args{"abpcplea", []string{"ale", "apple", "monkey", "plea"}},
			want: "apple",
		},
		{
			name: `s = "abpcplea", dictionary = ["a","b","c"]`,
			args: args{"abpcplea", []string{"a", "b", "c"}},
			want: "a",
		},
		{
			name: `s = "abpcplea", dictionary = ["z"]`,
			args: args{"abpcplea", []string{"z"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
