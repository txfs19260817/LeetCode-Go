package _671_Minimum_Number_of_Removals_to_Make_Mountain_Array

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	ans, left, right := n, make([]int, n), make([]int, n)
	for i := 0; i < n; i++ { // left LIS
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				left[i] = max(left[i], left[j]+1)
			}
		}
	}
	for i := n - 1; i >= 0; i-- { // right LIS
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] {
				right[i] = max(right[i], right[j]+1)
			}
		}
	}
	for i := 1; i < n-1; i++ {
		if left[i] == 0 || right[i] == 0 { // it is not a mountain if it is a monotonically inc/desc array
			continue
		}
		ans = min(ans, n-left[i]-right[i]-1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
