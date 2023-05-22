package leetcode

import "sort"

func numRescueBoats(people []int, limit int) int {
	ans, l, r := 0, 0, len(people)-1
	sort.Ints(people)
	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		ans++
	}
	return ans
}
