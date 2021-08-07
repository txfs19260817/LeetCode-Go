package _342_Power_of_Four

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}

func isPowerOfFour1(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
