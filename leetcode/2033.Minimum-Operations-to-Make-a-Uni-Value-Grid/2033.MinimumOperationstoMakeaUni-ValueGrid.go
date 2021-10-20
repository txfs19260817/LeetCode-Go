package _033_Minimum_Operations_to_Make_a_Uni_Value_Grid

import "sort"

func minOperations(grid [][]int, x int) int {
	ans, nums := 0, make([]int, 0, len(grid)*len(grid[0]))
	for _, line := range grid { // flatten
		nums = append(nums, line...)
	}
	sort.Ints(nums) // get mid
	mid := nums[len(nums)/2]
	for _, num := range nums {
		if (mid-num)%x != 0 { // infeasible if delta is not an integer multiple if x
			return -1
		}
		ans += abs(num - mid)
	}
	return ans / x
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
