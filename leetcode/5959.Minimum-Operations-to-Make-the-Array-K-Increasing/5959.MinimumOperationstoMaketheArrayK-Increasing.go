package leetcode

import "sort"

func kIncreasing(arr []int, k int) int {
	var keep int
	for i := 0; i < k; i++ {
		var lnds []int // longest non-descending seq
		for j := i; j < len(arr); j += k {
			idx := sort.SearchInts(lnds, arr[j]+1)
			if idx == len(lnds) {
				lnds = append(lnds, arr[j])
			} else {
				lnds[idx] = arr[j]
			}
		}
		keep += len(lnds)
	}
	return len(arr) - keep
}
