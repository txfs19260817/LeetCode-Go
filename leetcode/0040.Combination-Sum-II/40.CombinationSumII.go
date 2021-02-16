package _040_Combination_Sum_II

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	sort.Ints(candidates)
	var dfs func(candidates, path []int, target, index int)
	dfs = func(candidates, path []int, target, index int) {
		if target <= 0 {
			if target == 0 {
				temp := make([]int, len(path))
				copy(temp, path)
				ans = append(ans, temp)
			}
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			dfs(candidates, path, target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(candidates, []int{}, target, 0)
	return ans
}
