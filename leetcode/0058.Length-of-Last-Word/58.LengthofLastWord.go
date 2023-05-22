package leetcode

func lengthOfLastWord(s string) int {
	var l, r int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			r = i + 1
			break
		}
	}
	for i := r - 1; i >= 1; i-- {
		if s[i-1] == ' ' {
			l = i
			break
		}
	}
	return r - l
}
