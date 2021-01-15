package _057_Insert_Interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	// insertion
	insertionIdx := 0
	for i, interval := range intervals {
		if interval[0] >= newInterval[0] {
			break
		} else {
			insertionIdx = i + 1
		}
	}
	if insertionIdx == len(intervals) {
		intervals = append(intervals, newInterval)
	} else {
		intervals = append(intervals[:insertionIdx+1], intervals[insertionIdx:]...)
		intervals[insertionIdx] = newInterval
	}

	// merge
	res := make([][]int, 1, len(intervals))
	res[0] = intervals[0]
	intervals = intervals[1:]

	for _, interval := range intervals {
		if interval[0] > res[len(res)-1][1] {
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
