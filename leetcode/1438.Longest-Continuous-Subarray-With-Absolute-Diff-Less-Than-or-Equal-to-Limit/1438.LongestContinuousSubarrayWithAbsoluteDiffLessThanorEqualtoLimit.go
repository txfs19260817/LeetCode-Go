package leetcode

func longestSubarray(nums []int, limit int) int {
	var ans, l int
	var minQ, maxQ []int // dec, inc queues
	for r, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)

		// shrink window when exceed limit
		// the first element in minQ/maxQ is the smallest/largest element in the window
		for ; len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit; l++ {
			if minQ[0] == nums[l] {
				minQ = minQ[1:]
			}
			if maxQ[0] == nums[l] {
				maxQ = maxQ[1:]
			}
		}

		// update ans with current window size
		if length := r - l + 1; length > ans {
			ans = length
		}
	}
	return ans
}
