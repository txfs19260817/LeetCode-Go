package _894_Find_the_Student_that_Will_Replace_the_Chalk

import "sort"

func chalkReplacer(chalk []int, k int) int {
	var ans, sum int
	for _, c := range chalk {
		sum += c
	}
	k %= sum
	for i, c := range chalk {
		k -= c
		if k < 0 {
			ans = i
			break
		}
	}
	return ans
}

func chalkReplacer2(chalk []int, k int) int {
	for i := 1; i < len(chalk); i++ {
		chalk[i] += chalk[i-1]
	}
	k %= chalk[len(chalk)-1]
	idx := sort.SearchInts(chalk, k)
	if k == chalk[idx] { // upper bound
		idx++
	}
	return idx
}
