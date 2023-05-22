package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	ans, prefix := make([]int, len(queries)), make([]int, len(arr))
	prefix[0] = arr[0]
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] ^ arr[i]
	}
	for i, query := range queries {
		if query[0] == 0 {
			ans[i] = prefix[query[1]]
		} else {
			ans[i] = prefix[query[0]-1] ^ prefix[query[1]]
		}
	}
	return ans
}
