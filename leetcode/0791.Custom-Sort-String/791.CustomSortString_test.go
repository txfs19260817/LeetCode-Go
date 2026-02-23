package leetcode

import "testing"

func Test_customSortString(t *testing.T) {
	type args struct {
		order string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `order = "cbd", s = "abcd"`,
			args: args{"cbd", "abcd"},
			want: "",
		},
		{
			name: `order = "cbafg", s = "abcd"`,
			args: args{"cbafg", "abcd"},
			want: "",
		},
		{
			name: `order = "kqep", s = "pekeq"`,
			args: args{"kqep", "pekeq"},
			want: "kqeep",
		},
		{
			name: `order = "cba", s = "aaabbbccc"`,
			args: args{"cba", "aaabbbccc"},
			want: "cccbbbaaa",
		},
		{
			name: `order = "xyz", s = "abcd"`,
			args: args{"xyz", "abcd"},
			want: "",
		},
		{
			name: `order = "", s = "leetcode"`,
			args: args{"", "leetcode"},
			want: "",
		},
		{
			name: `order = "abc", s = ""`,
			args: args{"abc", ""},
			want: "",
		},
		{
			name: `order = "abc", s = "aaabccab"`,
			args: args{"abc", "aaabccab"},
			want: "aaaabbcc",
		},
		{
			name: `order = "zyxwvutsrqponmlkjihgfedcba", s = "abcabc"`,
			args: args{"zyxwvutsrqponmlkjihgfedcba", "abcabc"},
			want: "ccbbaa",
		},
		{
			name: `order = "a", s = "aaaa"`,
			args: args{"a", "aaaa"},
			want: "aaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := customSortString(tt.args.order, tt.args.s)
			if !isValidCustomSortResult(tt.args.order, tt.args.s, got) {
				t.Errorf("customSortString() = %v, invalid for order=%q s=%q", got, tt.args.order, tt.args.s)
			}
			if tt.want != "" && got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}

			got = customSortString2(tt.args.order, tt.args.s)
			if !isValidCustomSortResult(tt.args.order, tt.args.s, got) {
				t.Errorf("customSortString2() = %v, invalid for order=%q s=%q", got, tt.args.order, tt.args.s)
			}
			if tt.want != "" && got != tt.want {
				t.Errorf("customSortString2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isValidCustomSortResult(order, s, got string) bool {
	if len(s) != len(got) {
		return false
	}

	var sFreq, gotFreq [26]int
	for i := 0; i < len(s); i++ {
		sFreq[s[i]-'a']++
		gotFreq[got[i]-'a']++
	}
	if sFreq != gotFreq {
		return false
	}

	rank := make([]int, 26)
	for i := range rank {
		rank[i] = -1
	}
	for i := 0; i < len(order); i++ {
		rank[order[i]-'a'] = i
	}

	prevRank := -1
	for i := 0; i < len(got); i++ {
		r := rank[got[i]-'a']
		if r == -1 {
			continue
		}
		if r < prevRank {
			return false
		}
		prevRank = r
	}
	return true
}
