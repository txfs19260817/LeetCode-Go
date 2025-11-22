package leetcode

import (
	"maps"
	"slices"
)

func maximumTotalDamage(power []int) int64 {
	vec := map[int]int{}
	for _, x := range power {
		vec[x] += x
	}
	return int64(rob(vec))
}

func rob(vec map[int]int) int {
	sortedKeys := slices.Sorted(maps.Keys(vec))
	f := make([]int, len(vec)+1)
	j := 0
	for i, x := range sortedKeys {
		for sortedKeys[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+vec[x])
	}
	return f[len(vec)]
}
