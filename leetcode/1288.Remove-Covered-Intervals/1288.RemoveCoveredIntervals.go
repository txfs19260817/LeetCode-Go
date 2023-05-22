package leetcode

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	stack := make([][]int, 1, len(intervals))
	stack[0] = intervals[0]
	for _, interval := range intervals[1:] {
		if interval[1] > stack[len(stack)-1][1] {
			if interval[0] == stack[len(stack)-1][0] {
				stack[len(stack)-1][1] = interval[1]
			} else {
				stack = append(stack, interval)
			}
		}
	}
	return len(stack)
}
