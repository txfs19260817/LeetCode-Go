package leetcode

func brokenCalc(startValue int, target int) int {
	var cnt int
	for ; startValue < target; cnt++ {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
	}
	return cnt + startValue - target
}
