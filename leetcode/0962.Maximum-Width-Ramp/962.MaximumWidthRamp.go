package leetcode

func maxWidthRamp(nums []int) int {
	var ans int
	var sIdx []int // decreasing
	for i, num := range nums {
		if len(sIdx) == 0 || num < nums[sIdx[len(sIdx)-1]] {
			sIdx = append(sIdx, i)
		}
	}
	for j := len(nums) - 1; j >= 0; j-- {
		for len(sIdx) > 0 && nums[sIdx[len(sIdx)-1]] <= nums[j] {
			i := sIdx[len(sIdx)-1]
			sIdx = sIdx[:len(sIdx)-1]
			ans = max(ans, j-i)
		}
	}

	return ans
}
