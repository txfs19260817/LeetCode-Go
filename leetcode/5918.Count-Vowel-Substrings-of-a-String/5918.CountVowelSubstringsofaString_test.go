package _918_Count_Vowel_Substrings_of_a_String

import "testing"

func Test_countVowelSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "aeiouu",
			args: args{"aeiouu"},
			want: 2,
		},
		{
			name: "ughspuuoaaaoieiuiaoiuee",
			args: args{"ughspuuoaaaoieiuiaoiuee"},
			want: 76,
		},
		{
			name: "unicornarihan",
			args: args{"unicornarihan"},
			want: 0,
		},
		{
			name: "aeoibsddaaeiouudb",
			args: args{"aeoibsddaaeiouudb"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelSubstrings(tt.args.word); got != tt.want {
				t.Errorf("countVowelSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
