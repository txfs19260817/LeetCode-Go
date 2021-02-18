package _090_Subsets_II

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		ans = append(ans, combine(nums, k)...)
	}
	return ans
}

func combine(nums []int, k int) [][]int {
	var ans [][]int
	var dfs func(nums, path []int, k, index int)
	dfs = func(nums, path []int, k, index int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := index; i < len(nums); i++ {
			if i > index && nums[i-1] == nums[i] {
				continue
			}
			path = append(path, nums[i])
			dfs(nums, path, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, []int{}, k, 0)
	return ans
}
