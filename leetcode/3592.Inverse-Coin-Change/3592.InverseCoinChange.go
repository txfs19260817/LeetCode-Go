package leetcode

// 性质 1：金额 i 不受大于 i 的硬币影响
// 性质 2：在处理 i 之前，所有小于 i 的硬币状态都已经唯一确定
// 性质 3：面值 i 的硬币对金额 i 的贡献恰好只有 1
func findCoins(numWays []int) []int {
	// numWays[i-1] 表示凑出金额 i 的方案数
	var ans []int // 硬币集合
	n := len(numWays)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		numWay := numWays[i-1]
		// 真实值和“已有小硬币”算出来的值完全一致。
		if numWay == dp[i] {
			continue
		}
		// 应该多了恰好 1 种方案 [i]，所以说明面值 i 的硬币存在
		if numWay != dp[i]+1 {
			return nil
		}
		// 确认面值 i 存在，加入答案
		ans = append(ans, i)
		// 把这个新硬币用coin change 2的完全背包方法加入计算
		for j := i; j <= n; j++ {
			dp[j] += dp[j-i]
		}
	}
	// 此时 dp[1:] == numWays
	return ans
}
