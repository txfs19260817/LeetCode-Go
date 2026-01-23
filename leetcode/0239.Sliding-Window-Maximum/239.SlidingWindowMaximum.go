package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	ans, maxQ := make([]int, 0, len(nums)-k+1), make([]int, 0, len(nums)-k+1)
	for i, x := range nums {
		// enqueue
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		// fill the window
		if i < k-1 {
			continue
		}

		ans = append(ans, nums[maxQ[0]])

		// dequeue
		if len(maxQ) > 0 && maxQ[0] <= i-k+1 {
			maxQ = maxQ[1:]
		}
	}

	return ans
}
