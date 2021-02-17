package _046_Permutations

func permute(nums []int) [][]int {
	var ans [][]int
	var dfs func([]int, int)
	dfs = func(nums []int, index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
			return
		}
		for i := index; i < len(nums); i++ {
			nums[index], nums[i] = nums[i], nums[index]
			dfs(nums, index+1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	dfs(nums, 0)
	return ans
}
