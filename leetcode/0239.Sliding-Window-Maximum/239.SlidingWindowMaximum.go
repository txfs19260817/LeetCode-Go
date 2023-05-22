package leetcode

type pair struct {
	idx, val int
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = -1 << 30
	}
	var idxQueue []int
	updateQueue := func(curV, curI, startIdx int) {
		if len(idxQueue) > 0 && idxQueue[0] <= startIdx {
			idxQueue = idxQueue[1:]
		}
		for len(idxQueue) > 0 && nums[idxQueue[len(idxQueue)-1]] <= curV {
			idxQueue = idxQueue[:len(idxQueue)-1]
		}
		idxQueue = append(idxQueue, curI)
	}
	for l, r := 0, 0; r < len(nums); {
		updateQueue(nums[r], r, l)
		r++
		if r-l > k {
			l++
		}
		ans[l] = max(ans[l], nums[idxQueue[0]])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
