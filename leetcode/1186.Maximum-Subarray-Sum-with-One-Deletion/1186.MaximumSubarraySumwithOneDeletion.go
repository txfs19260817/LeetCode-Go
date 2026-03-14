package leetcode

import (
	"slices"
)

func maximumSum(arr []int) int {
	ans := slices.Min(arr)
	f0, f1 := ans, ans // f0: no delete; f1: delete arr[i-1] and keep arr[i]
	for _, x := range arr {
		f1 = max(f1+x, f0)  // already deleted or keep all except current one
		f0 = max(f0, 0) + x // max sub array (delete nothing)
		ans = max(ans, f1, f0)
	}
	return ans
}
