package leetcode

import "sort"

func findWinners(matches [][]int) [][]int {
	w, l := map[int]int{}, map[int]int{}
	var gw, gl []int
	for _, match := range matches {
		w[match[0]]++
		l[match[1]]++
	}
	for p := range w {
		if l[p] == 0 {
			gw = append(gw, p)
		}
	}
	for p, g := range l {
		if g == 1 {
			gl = append(gl, p)
		}
	}
	sort.Ints(gw)
	sort.Ints(gl)
	return [][]int{gw, gl}
}
