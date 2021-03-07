package _647_Palindromic_Substrings

func countSubstrings(s string) int {
	var ans int
	for i := range s {
		ans += expansion(s, i, i) + expansion(s, i, i+1)
	}
	return ans
}

func expansion(s string, l, r int) int {
	var count int
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		count++
	}
	return count
}
