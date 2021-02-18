package _077_Combinations

func combine(n int, k int) [][]int {
	var ans [][]int
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	var dfs func(nums, path []int, k, index int)
	dfs = func(nums, path []int, k, index int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, path, k, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, []int{}, k, 0)
	return ans
}
