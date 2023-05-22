package leetcode

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}
	return ans
}
