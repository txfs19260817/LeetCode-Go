package leetcode

func isSubsequence(s string, t string) bool {
	for i, j := 0, 0; i < len(s); i, j = i+1, j+1 {
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j >= len(t) {
			return false
		}
	}
	return true
}
