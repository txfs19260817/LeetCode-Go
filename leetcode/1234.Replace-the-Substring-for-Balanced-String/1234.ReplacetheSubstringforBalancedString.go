package leetcode

func balancedString(s string) int {
	count, k, l, r, ans := map[byte]int{}, len(s)/4, 0, 0, len(s)
	for _, v := range s {
		count[byte(v)]++
	}
	for l < len(s) {
		if r < len(s) && (count['Q'] > k || count['W'] > k || count['E'] > k || count['R'] > k) {
			count[s[r]]--
			r++
		} else {
			count[s[l]]++
			l++
		}
		if !(count['Q'] > k || count['W'] > k || count['E'] > k || count['R'] > k) && ans > r-l {
			ans = r - l
		}
	}
	return ans
}
