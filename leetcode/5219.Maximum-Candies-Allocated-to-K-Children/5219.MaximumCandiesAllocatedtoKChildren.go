package _219_Maximum_Candies_Allocated_to_K_Children

func maximumCandies(candies []int, k int64) int {
	var l, r int
	for _, candy := range candies {
		if candy > r {
			r = candy
		}
	}
	for l+1 < r {
		mid := l + (r-l)/2
		if check(candies, k, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	if check(candies, k, r) {
		return r
	}
	return l
}

func check(candies []int, k int64, target int) bool {
	var piles int
	for _, candy := range candies {
		piles += candy / target
		if int64(piles) >= k {
			return true
		}
	}
	return false
}
