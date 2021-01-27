package _125_Valid_Palindrome

import "strings"

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !isAlphanumeric(s[l]) {
			l++
		}
		for l < r && !isAlphanumeric(s[r]) {
			r--
		}
		if !isEqual(s[l], s[r]) {
			return false
		}
	}
	return true
}

func isEqual(a, b byte) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b))
}

func isAlphanumeric(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}
