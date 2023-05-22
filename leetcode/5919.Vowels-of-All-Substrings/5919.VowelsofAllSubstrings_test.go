package leetcode

import "testing"

func Test_countVowels(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "aba",
			args: args{"aba"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowels(tt.args.word); got != tt.want {
				t.Errorf("countVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
