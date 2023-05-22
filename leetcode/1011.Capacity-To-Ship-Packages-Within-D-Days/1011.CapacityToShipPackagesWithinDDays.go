package leetcode

func shipWithinDays(weights []int, days int) int {
	daysRequired := func(capacity int) int {
		sum, d := 0, 1
		for _, weight := range weights {
			if sum+weight <= capacity {
				sum += weight
			} else {
				d++
				sum = weight
			}
		}
		return d
	}
	var l, r int // l = max(weights), r = sum(weights)
	for _, weight := range weights {
		if l < weight {
			l = weight
		}
		r += weight
	}
	for l < r {
		if mid := l + (r-l)/2; daysRequired(mid) > days {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
