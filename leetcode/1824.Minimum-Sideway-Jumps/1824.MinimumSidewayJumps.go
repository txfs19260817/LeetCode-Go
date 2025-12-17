package leetcode

import "slices"

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	d := [3]int{1, 0, 1}
	for _, o := range obstacles[1:] {
		for i := range d {
			if o-1 == i {
				d[i] = n
			}
		}
		minCnt := slices.Min(d[:])
		for i := range d {
			if o-1 != i {
				d[i] = min(d[i], minCnt+1)
			}
		}
	}
	return slices.Min(d[:])
}
