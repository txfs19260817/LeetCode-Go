package _567_Maximum_Length_of_Subarray_With_Positive_Product

func getMaxLen(nums []int) int {
	var ans int
	// split nums into sub-arrays by 0
	var subArrays [][]int
	var sub []int
	for _, num := range nums {
		if num == 0 {
			subArrays = append(subArrays, sub)
			sub = nil
		} else {
			sub = append(sub, num)
		}
	}
	subArrays = append(subArrays, sub)
	for _, arr := range subArrays {
		if len(arr) <= ans {
			continue
		}
		cntNeg, firstNegIdx, lastNegIdx := 0, -1, 0
		for i, a := range arr {
			if a < 0 {
				cntNeg++
				lastNegIdx = i
				if firstNegIdx == -1 {
					firstNegIdx = i
				}
			}
		}
		if cntNeg%2 == 0 { // if the number of negative number is even
			ans = max(ans, len(arr))
		} else { // else, max([:lastNegIdx), (firstNegIdx:])
			ans = max(ans, max(lastNegIdx, len(arr)-firstNegIdx-1))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
