package leetcode

import "math"

func maxPoints(points [][]int) int {
	var ans int
	for i, p := range points {
		k2cnt := map[float64]int{}
		for _, q := range points[i+1:] {
			dx, dy := q[0]-p[0], q[1]-p[1]
			k := math.MaxFloat64
			if dx != 0 {
				k = float64(dy) / float64(dx)
			}
			k2cnt[k]++
			ans = max(ans, k2cnt[k])
		}
	}
	return ans + 1 // p itself
}
