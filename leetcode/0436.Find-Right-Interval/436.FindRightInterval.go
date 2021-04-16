package _436_Find_Right_Interval

import "sort"

func findRightInterval(intervals [][]int) []int {
	ans, starts, start2idx := make([]int, len(intervals)), make([]int, 0, len(intervals)), map[int]int{}
	// build a sorted array for start indices & a map from start to its index in intervals
	for i, interval := range intervals {
		if insertIdx := sort.SearchInts(starts, interval[0]); insertIdx == len(starts) {
			starts = append(starts, interval[0])
		} else {
			starts = append(starts[:insertIdx+1], starts[insertIdx:]...)
			starts[insertIdx] = interval[0]
		}
		start2idx[interval[0]] = i
	}
	// find and assign index to ans
	for i, interval := range intervals {
		if startIdx := sort.SearchInts(starts, interval[1]); startIdx == len(starts) {
			ans[i] = -1
		} else {
			ans[i] = start2idx[starts[startIdx]]
		}
	}
	return ans
}
