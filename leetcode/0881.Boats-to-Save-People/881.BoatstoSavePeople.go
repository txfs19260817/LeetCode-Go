package _881_Boats_to_Save_People

import "sort"

func numRescueBoats(people []int, limit int) int {
	ans, l, r := 0, 0, len(people)-1
	sort.Ints(people)
	for l < r {
		if people[l]+people[r] <= limit {
			l++
		}
		ans++
		r--
	}
	if l == r {
		ans++
	}
	return ans
}
