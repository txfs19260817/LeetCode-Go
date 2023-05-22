package leetcode

func rearrangeArray(nums []int) []int {
	ans, i, j := make([]int, len(nums)), 0, 1
	for _, num := range nums {
		if num > 0 {
			ans[i] = num
			i += 2
		} else {
			ans[j] = num
			j += 2
		}
	}
	return ans
}
