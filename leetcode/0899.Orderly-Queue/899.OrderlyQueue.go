package leetcode

import "sort"

func orderlyQueue(s string, k int) string {
	if k == 1 { // necklace
		ans := s
		for i := 0; i < len(s); i++ {
			s = string(append([]byte(s[1:]), s[0]))
			if ans > s {
				ans = s
			}
		}
		return ans
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
