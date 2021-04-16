package _435_Non_overlapping_Intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}
	var ans int
	sort.Slice(intervals, func(i, j int) bool { // sort by the end of each interval
		return intervals[i][1] < intervals[j][1]
	})
	for preEnd, i := intervals[0][1], 1; i < len(intervals); i++ {
		if preEnd > intervals[i][0] { // overlapped, remove the current interval
			ans++
		} else { // non-overlapped, update last end
			preEnd = intervals[i][1]
		}
	}
	return ans
}
