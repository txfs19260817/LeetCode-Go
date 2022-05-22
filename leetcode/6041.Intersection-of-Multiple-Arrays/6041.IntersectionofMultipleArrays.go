package _041_Intersection_of_Multiple_Arrays

import "sort"

func intersection(nums [][]int) []int {
	var ans []int
	cnt := map[int]int{}
	for _, arr := range nums {
		for _, v := range arr {
			cnt[v]++
		}
	}
	for v, c := range cnt {
		if c == len(nums) {
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return ans
}
