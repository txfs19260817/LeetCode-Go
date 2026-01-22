package leetcode

func longestSubarray(nums []int, limit int) int {
	var ans, l int
	var minQ, maxQ []int // inc/dec index queues
	for i, x := range nums {
		// right insert
		for len(minQ) > 0 && x <= nums[minQ[len(minQ)-1]] { // minQ is a monotonic increasing queue
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, i)
		for len(maxQ) > 0 && x >= nums[maxQ[len(maxQ)-1]] { // maxQ is a monotonic decreasing queue
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, i)

		// left pop
		for nums[maxQ[0]]-nums[minQ[0]] > limit {
			l++
			if l > maxQ[0] {
				maxQ = maxQ[1:]
			}
			if l > minQ[0] {
				minQ = minQ[1:]
			}
		}

		// update ans
		ans = max(ans, i-l+1)
	}
	return ans
}
