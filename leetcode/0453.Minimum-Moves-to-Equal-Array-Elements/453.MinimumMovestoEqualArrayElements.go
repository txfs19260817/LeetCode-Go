package leetcode

// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/solution/geekplayers-leetcode-ac-qing-xi-yi-dong-4ei1r/
func minMoves(nums []int) int {
	ans, minV := 0, nums[0]
	for _, n := range nums[1:] {
		if minV > n {
			minV = n
		}
	}
	for _, n := range nums {
		ans += n - minV
	}
	return ans
}
