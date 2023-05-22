package leetcode

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	dfs(&ans, []int{}, nums, 0, -1<<31)
	return ans
}

func dfs(ans *[][]int, path, nums []int, index, lastValue int) {
	if index == len(nums) {
		if len(path) >= 2 {
			temp := make([]int, len(path))
			copy(temp, path)
			*ans = append(*ans, temp)
		}
		return
	}
	if nums[index] >= lastValue {
		path = append(path, nums[index])
		dfs(ans, path, nums, index+1, nums[index])
		path = path[:len(path)-1]
	}
	if nums[index] != lastValue {
		dfs(ans, path, nums, index+1, lastValue)
	}
}
