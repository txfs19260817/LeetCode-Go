package _016_3Sum_Closest

import "sort"

func threeSumClosest(nums []int, target int) int {
	diff, ans := 2147483647, 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < diff {
				diff, ans = abs(sum-target), sum
			}
			if diff == 0 {
				return ans
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
