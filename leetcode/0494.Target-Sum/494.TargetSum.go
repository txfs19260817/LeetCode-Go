package leetcode

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

func findTargetSumWays1(nums []int, target int) int {
	var ans int
	var dfs func(idx, sum int)
	dfs = func(idx, sum int) {
		if idx == len(nums) {
			if target == sum {
				ans++
			}
			return
		}
		dfs(idx+1, sum+nums[idx])
		dfs(idx+1, sum-nums[idx])
	}
	dfs(0, 0)
	return ans
}
