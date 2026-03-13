package leetcode

func isArraySpecial(nums []int, queries [][]int) []bool {
	prefix := make([]int, len(nums)+1)
	for i := 1; i < len(prefix); i++ {
		bad := 0
		if i < len(nums) && nums[i-1]%2 == nums[i]%2 {
			bad = 1
		}
		prefix[i] = prefix[i-1] + bad
	}

	ans := make([]bool, len(queries))
	for i, query := range queries {
		ans[i] = prefix[query[1]]-prefix[query[0]] == 0
	}
	return ans
}
