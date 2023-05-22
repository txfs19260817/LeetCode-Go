package leetcode

func maximumTop(nums []int, k int) int {
	N := len(nums)
	// if no moves allowed, return the topmost element if any
	if k == 0 {
		if N >= 1 {
			return nums[0]
		}
		return -1
	}
	// if only one move is allowed, we can only remove the topmost element
	if k == 1 {
		if N == 1 {
			return -1
		}
		return nums[1]
	}
	// if `N == 1`, we can return the topmost element if `k` is a even number (keep removing the topmost element and adding it back).
	if N == 1 {
		if k%2 == 0 {
			return nums[0]
		}
		return -1
	}
	// we can take `min(k-1, N)` elements and put back the largest one on the top
	mx := nums[0]
	for i := 1; i < min(k-1, N); i++ {
		mx = max(mx, nums[i])
	}
	// If `k < N`, we can take all the topmost `k` elements and return the one left at the top
	if k < N {
		mx = max(mx, nums[k])
	}
	return mx
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
