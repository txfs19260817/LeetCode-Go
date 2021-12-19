package _481_Least_Number_of_Unique_Integers_after_K_Removals

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	type pair struct {
		val, cnt int
	}
	v2c := map[int]int{}
	for _, v := range arr {
		v2c[v]++
	}
	pairs := make([]pair, 0, len(v2c))
	for v, c := range v2c {
		pairs = append(pairs, pair{val: v, cnt: c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt < pairs[j].cnt })
	for _, p := range pairs {
		if p.cnt > k {
			break
		}
		k -= p.cnt
		delete(v2c, p.val)
	}
	return len(v2c)
}
