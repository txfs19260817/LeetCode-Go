package leetcode

func kthFactor(n int, k int) int {
	i := 1
	for ; i*i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	i--           // compensate the last i++
	if i*i == n { // remove duplicated factors
		i--
	}
	for ; i > 0; i-- {
		if n%i == 0 {
			k--
			if k == 0 { // the latter half of the factors in asc
				return n / i
			}
		}
	}
	return -1
}
