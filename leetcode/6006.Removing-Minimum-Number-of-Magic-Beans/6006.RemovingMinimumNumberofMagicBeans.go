package _006_Removing_Minimum_Number_of_Magic_Beans

import "sort"

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	presum := make([]int64, n)
	presum[0] = int64(beans[0])
	for i := 1; i < n; i++ {
		presum[i] = presum[i-1] + int64(beans[i])
	}
	sum := presum[n-1]
	cnt := sum - presum[0] - int64(n-1)*int64(beans[0])
	for i := 1; i < n; i++ {
		if newCnt := sum - presum[i] - int64(n-i-1)*int64(beans[i]) + presum[i-1]; newCnt < cnt {
			cnt = newCnt
		}
	}
	return cnt
}
