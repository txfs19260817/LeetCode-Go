package leetcode

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var ans int
	var deltas []int
	for i, c := range capacity {
		r := rocks[i]
		if c == r {
			ans++
			continue
		}
		deltas = append(deltas, c-r)
	}
	sort.Ints(deltas)
	for _, d := range deltas {
		if additionalRocks < d {
			break
		}
		additionalRocks -= d
		ans++
	}
	return ans
}
