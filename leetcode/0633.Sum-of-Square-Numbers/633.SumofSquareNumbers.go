package leetcode

func judgeSquareSum(c int) bool {
	m := map[int]bool{}
	for i := 0; i*i <= c; i++ {
		m[i*i] = true
	}
	for n := range m {
		if m[c-n] {
			return true
		}
	}
	return false
}
