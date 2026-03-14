package leetcode

func pourWater(heights []int, volume int, k int) []int {
	for range volume {
		mn, imn := heights[k], k
		for i := k - 1; i >= 0; i-- {
			if heights[i] > mn {
				break
			}
			if heights[i] < mn {
				mn, imn = heights[i], i
			}
		}
		if mn < heights[k] {
			heights[imn]++
			continue
		}

		mn, imn = heights[k], k
		for i := k + 1; i < len(heights); i++ {
			if heights[i] > mn {
				break
			}
			if heights[i] < mn {
				mn, imn = heights[i], i
			}
		}
		if mn < heights[k] {
			heights[imn]++
			continue
		}

		heights[k]++
	}
	return heights
}
