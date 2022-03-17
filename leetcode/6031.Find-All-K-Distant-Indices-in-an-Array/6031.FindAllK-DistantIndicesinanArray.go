package _031_Find_All_K_Distant_Indices_in_an_Array

func findKDistantIndices(nums []int, key int, k int) []int {
	var ans []int
	var intervals [][]int
	for i, num := range nums {
		if num == key {
			intervals = append(intervals, []int{max(0, i-k), min(i+k, len(nums)-1)})
		}
	}

	res := make([][]int, 1, len(intervals))
	res[0] = intervals[0]
	intervals = intervals[1:]
	for _, interval := range intervals {
		if res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}

	for _, interval := range res {
		for i := interval[0]; i <= interval[1]; i++ {
			ans = append(ans, i)
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
