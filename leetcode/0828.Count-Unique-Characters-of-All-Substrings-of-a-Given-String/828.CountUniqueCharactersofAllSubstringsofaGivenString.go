package leetcode

func uniqueLetterString(s string) int {
	ans, c2i := 0, map[rune][]int{}
	for i, c := range s {
		c2i[c] = append(c2i[c], i)
	}
	for _, indices := range c2i {
		for i := 0; i < len(indices); i++ {
			prev, next := -1, len(s)
			if i > 0 {
				prev = indices[i-1]
			}
			if i < len(indices)-1 {
				next = indices[i+1]
			}
			ans += (indices[i] - prev) * (next - indices[i])
		}
	}
	return ans
}

func uniqueLetterString2(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		l := i - 1
		for l >= 0 && s[i] != s[l] {
			l--
		}
		r := i + 1
		for r < len(s) && s[i] != s[r] {
			r++
		}
		res += (i - l) * (r - i)
	}
	return res
}
