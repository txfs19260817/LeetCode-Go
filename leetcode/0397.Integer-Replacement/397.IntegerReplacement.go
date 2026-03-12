package leetcode

func integerReplacement(n int) int {
	var ans int
	for ; n > 1; ans++ {
		if n%2 == 0 {
			n >>= 1
		} else {
			// For odd n, prefer n-1 when it quickly creates more trailing zeros in binary.
			// n == 3 is a special case: 3 -> 2 -> 1 is shorter than 3 -> 4 -> 2 -> 1.
			if n == 3 || n%4 == 1 {
				n--
			} else {
				n++
			}
		}
	}
	return ans
}
