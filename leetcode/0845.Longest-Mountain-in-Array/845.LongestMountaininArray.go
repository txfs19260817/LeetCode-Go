package _845_Longest_Mountain_in_Array

func longestMountain(arr []int) int {
	l, r, res, isAscending := 0, 0, 0, true
	for l < len(arr) {
		if r+1 < len(arr) && (isAscending && arr[r] < arr[r+1] || l != r && arr[r] > arr[r+1]) {
			if arr[r] > arr[r+1] {
				isAscending = false
			}
			r++
		} else {
			if r != l && !isAscending {
				res = max(res, r-l+1)
			}
			l++
			r = max(r, l)
			if l == r {
				isAscending = true
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
