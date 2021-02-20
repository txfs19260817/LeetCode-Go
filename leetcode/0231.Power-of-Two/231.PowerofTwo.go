package _231_Power_of_Two

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfTwo1(n int) bool {
	var count int
	if n <= 0 {
		return false
	}
	for i := 0; i < 32; i++ {
		count += n & 1
		n >>= 1
	}
	return count == 1
}
