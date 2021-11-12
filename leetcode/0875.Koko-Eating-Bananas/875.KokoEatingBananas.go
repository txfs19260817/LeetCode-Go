package _875_Koko_Eating_Bananas

func minEatingSpeed(piles []int, h int) int {
	check := func(k int) bool {
		var time int
		for _, pile := range piles {
			time += (pile + k - 1) / k
		}
		return time <= h
	}
	l, r := 1, 1000000000
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
