package leetcode

import "strconv"

func nextGreaterElement(n int) int {
	digits := []byte(strconv.Itoa(n))
	i, j := len(digits)-2, len(digits)-1
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}
	if i == -1 {
		return -1
	}
	for digits[i] >= digits[j] {
		j--
	}
	digits[i], digits[j] = digits[j], digits[i]
	for l, r := i+1, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	ans, _ := strconv.Atoi(string(digits))
	if ans > 1<<31-1 {
		return -1
	}
	return ans
}
