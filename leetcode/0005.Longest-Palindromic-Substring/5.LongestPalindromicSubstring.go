package _005_Longest_Palindromic_Substring

func longestPalindrome(s string) string {
	var start, end int
	for i := range s {
		l1, r1 := expansion(s, i, i)
		l2, r2 := expansion(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func expansion(s string, l, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
