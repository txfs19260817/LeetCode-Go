package leetcode

import (
	"sort"
	"strconv"
)

func largestInteger(num int) int {
	s := strconv.Itoa(num)
	a := [2][]rune{}
	for _, v := range s {
		a[v%2] = append(a[v%2], v)
	}
	sort.Slice(a[0], func(i, j int) bool { return a[0][i] > a[0][j] })
	sort.Slice(a[1], func(i, j int) bool { return a[1][i] > a[1][j] })

	t := make([]rune, len(s))
	for i, ch := range s {
		j := ch % 2
		t[i] = a[j][0]
		a[j] = a[j][1:]
	}
	ans, _ := strconv.Atoi(string(t))
	return ans
}
