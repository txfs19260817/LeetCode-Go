package _066_Plus_One

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += c
		c = 0 // important
		if digits[i] > 9 {
			c = digits[i] / 10
			digits[i] %= 10
		}
	}
	if c != 0 {
		return append([]int{c}, digits...)
	}
	return digits
}
