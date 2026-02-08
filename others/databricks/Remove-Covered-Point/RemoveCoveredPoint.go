package databricks

// DeleteCoveredPoint removes the idx-th integer (0-based) from the flattened
// sequence of all integers covered by the given non-overlapping intervals
// [start, end). Returns the updated list of intervals.
func DeleteCoveredPoint(intervals [][]int, idx int) [][]int {
	remaining := idx
	for i, interval := range intervals {
		start, end := interval[0], interval[1]
		size := end - start
		if remaining < size {
			point := start + remaining
			if size == 1 {
				// Remove interval entirely
				result := make([][]int, 0, len(intervals)-1)
				result = append(result, intervals[:i]...)
				result = append(result, intervals[i+1:]...)
				return result
			}
			if point == start {
				// Shrink left
				intervals[i] = []int{start + 1, end}
				return intervals
			}
			if point == end-1 {
				// Shrink right
				intervals[i] = []int{start, end - 1}
				return intervals
			}
			// Split into two intervals
			result := make([][]int, 0, len(intervals)+1)
			result = append(result, intervals[:i]...)
			result = append(result, []int{start, point})
			result = append(result, []int{point + 1, end})
			result = append(result, intervals[i+1:]...)
			return result
		}
		remaining -= size
	}
	return intervals // idx out of range (should not happen with valid input)
}
