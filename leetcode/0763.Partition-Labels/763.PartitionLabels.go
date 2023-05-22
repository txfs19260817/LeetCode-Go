package leetcode

import "sort"

func partitionLabels(S string) []int {
	var ans []int
	last := map[byte]int{}
	for i, c := range S {
		last[byte(c)] = i
	}
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = last[S[start]]
		for i := start; i <= end; i++ {
			if end < last[S[i]] {
				end = last[S[i]]
			}
		}
		ans = append(ans, end-start+1)
	}
	return ans
}

func partitionLabels2(S string) []int {
	var ans []int
	first, last := map[rune]int{}, map[rune]int{}
	for i, c := range S {
		last[c] = i
		if _, ok := first[c]; !ok {
			first[c] = i
		}
	}
	intervals := make([][]int, 0, len(first))
	for k := range first {
		intervals = append(intervals, []int{first[k], last[k]})
	}
	intervals = merge(intervals)
	for _, interval := range intervals {
		ans = append(ans, interval[1]-interval[0]+1)
	}
	return ans
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 1, len(intervals))
	res[0] = intervals[0]
	for _, interval := range intervals[1:] {
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
