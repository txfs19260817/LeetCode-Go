package _047_Permutations_II

func permuteUnique(nums []int) [][]int {
	var ans [][]int
	var dfs func([]int, int)
	dfs = func(nums []int, index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
			return
		}
		m := map[int]bool{}
		for i := index; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			nums[index], nums[i] = nums[i], nums[index]
			dfs(nums, index+1)
			nums[index], nums[i] = nums[i], nums[index]
			m[nums[i]] = true
		}
	}
	dfs(nums, 0)
	return ans
}
