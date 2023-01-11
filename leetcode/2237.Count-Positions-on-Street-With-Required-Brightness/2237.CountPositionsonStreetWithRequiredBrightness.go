package _237_Count_Positions_on_Street_With_Required_Brightness

/*
	Range addition / sprint training

0. get intervals
1. create an 1-d array with required length
2. arr[l]++ ; arr[r+1]-- (r+1 < len(arr))
3. get prefix sum of arr
*/
func meetRequirement(n int, lights [][]int, requirement []int) int {
	ans, lamps := 0, make([]int, len(requirement))
	for _, light := range lights {
		position, length := light[0], light[1]
		l, r := max(0, position-length), min(n-1, position+length) // inclusive covered interval
		lamps[l]++
		if r+1 < len(lamps) {
			lamps[r+1]--
		}
	}
	for i := 1; i < len(lamps); i++ {
		lamps[i] += lamps[i-1]
	}
	for i, lamp := range lamps {
		if lamp >= requirement[i] {
			ans++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
