package _956_Find_First_Palindromic_String_in_the_Array

func firstPalindrome(words []string) string {
	ans, isPalindrome := "", func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}
	for _, word := range words {
		if isPalindrome(word) {
			ans = word
			break
		}
	}
	return ans
}
