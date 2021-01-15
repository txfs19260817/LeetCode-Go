package _056_Merge_Intervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 1, len(intervals))
	res[0] = intervals[0]
	intervals = intervals[1:]
	for _, interval := range intervals {
		if res[len(res)-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
