package leetcode

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	dict, window := map[rune]int{}, map[rune]int{}
	for _, c := range t {
		dict[c]++
	}

	res, l, r, flag := s, 0, 0, false
	for l < len(s)-len(t)+1 {
		if r < len(s) && !isMapContain(dict, window) {
			window[rune(s[r])]++
			r++
		} else {
			window[rune(s[l])]--
			l++
		}
		if r-l <= len(res) && isMapContain(dict, window) { // `=` is suitable for res==s
			flag = true
			res = s[l:r]
		}
	}
	if !flag {
		return ""
	}
	return res
}

// isMapContain returns true if b>=a
func isMapContain(a, b map[rune]int) bool {
	for k, v := range a {
		if b[k] < v {
			return false
		}
	}
	return true
}
