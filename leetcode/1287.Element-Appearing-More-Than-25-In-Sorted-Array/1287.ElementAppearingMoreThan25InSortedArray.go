package _287_Element_Appearing_More_Than_25_In_Sorted_Array

func findSpecialInteger(arr []int) int {
	// If a value occurrs more than 25% of the time in a sorted array, then it must be on at least one of these three positions.
	for _, i := range []int{len(arr) / 4, len(arr) / 2, len(arr) * 3 / 4} {
		// A value occurrs more than 25% of the time if it occurrs >= Math.floor(n/4) + 1 times. Integer division does the .floor() for us.
		if f := findFirstOccurrence(arr, arr[i]); arr[f] == arr[f+len(arr)/4] {
			return arr[f]
		}
	}
	return -1
}

func findFirstOccurrence(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		if mid := l + (r-l)/2; arr[mid] >= target {
			r = mid
		} else {
			l = mid
		}
	}
	if arr[l] == target {
		return l
	}
	return r
}
