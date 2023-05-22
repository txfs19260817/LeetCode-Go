package leetcode

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) - 1) / 2
	for _, n := range nums {
		sum -= n
	}
	return sum
}

func missingNumber1(nums []int) int {
	xors, div, a, b, found := 0, 1, 0, 0, false
	for i, n := range nums {
		xors ^= n ^ i
	}
	if xors == 0 {
		return len(nums)
	}
	for xors&div == 0 {
		div <<= 1
	}
	for i, n := range nums {
		if div&n == 0 {
			a ^= n
		} else {
			b ^= n
		}
		if div&i == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}
	for _, n := range nums {
		if n == a {
			found = true
			break
		}
	}
	if found {
		return b
	}
	return a
}
