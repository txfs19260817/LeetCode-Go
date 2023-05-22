package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	m := map[byte]int{} //char:freq
	l, r, ans := 0, 0, 1
	for l < len(s) {
		if r < len(s) && m[s[r]] == 0 {
			m[s[r]]++
			r++
		} else {
			m[s[l]]--
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	return ans
}
