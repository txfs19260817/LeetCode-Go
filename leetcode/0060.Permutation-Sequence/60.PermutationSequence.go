package _060_Permutation_Sequence

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	permutation, token := make([]int, n+1), make([]int, n) // permutation = n!; token = [1,2,3,...,n]
	permutation[0] = 1
	for i := 1; i < n+1; i++ {
		permutation[i] = i * permutation[i-1]
		token[i-1] = i
	}
	k--

	ans := make([]byte, 0, n)
	for n > 0 {
		n-- // use n-1
		i := k / permutation[n]
		k %= permutation[n]
		ans = strconv.AppendInt(ans, int64(token[i]), 10)
		token = append(token[:i], token[i+1:]...) // remove ith num from token
	}
	return string(ans)
}
