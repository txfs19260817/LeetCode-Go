package k

import "sort"

func spareTime(meetings [][][]int) (ans [][]int) {
	merged := meetings[0]
	for _, m := range meetings[1:] {
		merged = mergeIntervals(merged, m)
	}
	if 0 < merged[0][0] {
		ans = append(ans, []int{0, merged[0][0]})
	}
	for i := 0; i < len(merged)-1; i++ {
		if merged[i][1] < merged[i+1][0] {
			ans = append(ans, []int{merged[i][1], merged[i+1][0]})
		}
	}
	if merged[len(merged)-1][1] < 2400 {
		ans = append(ans, []int{merged[len(merged)-1][1], 2400})
	}
	return
}

func mergeIntervals(a, b [][]int) [][]int {
	ans, intervals := make([][]int, 0, len(a)+len(b)), append(a, b...)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	ans = append(ans, intervals[0])
	for _, interval := range intervals[1:] {
		if interval[0] > ans[len(ans)-1][1] {
			ans = append(ans, interval)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
