package _891_Cutting_Ribbons

func maxLength(ribbons []int, k int) int {
	l, r := 0, 100000
	for l < r {
		mid, cnt := l+(r-l+1)/2, 0
		for _, ribbon := range ribbons {
			cnt += ribbon / mid
		}
		if cnt >= k {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
