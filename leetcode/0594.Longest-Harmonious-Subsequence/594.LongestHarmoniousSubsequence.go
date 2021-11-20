package _594_Longest_Harmonious_Subsequence

import "sort"

func findLHS(nums []int) int {
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		j := i
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
		}
		nextMove := j
		for ; j < len(nums) && nums[j]-1 == nums[i]; j++ {
			if l := j - i + 1; l > ans {
				ans = l
			}
		}
		i = nextMove
	}
	return ans
}

func findLHS2(nums []int) int {
	ans, cnt := 0, map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c1 := cnt[num+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}
	return ans
}
