package _991_Broken_Calculator

func brokenCalc(startValue int, target int) int {
	var ans int
	for startValue < target {
		ans++
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
	}
	return ans + startValue - target
}
