package _018_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i-1] == nums[i] || nums[i]+nums[n-1]+nums[n-2]+nums[n-3] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j-1] == nums[j] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for l, r := j+1, n-1; l < r; {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					r--
					l++
					for l < r && nums[r+1] == nums[r] {
						r--
					}
					for l < r && nums[l-1] == nums[l] {
						l++
					}
				}
			}
		}
	}
	return res
}
