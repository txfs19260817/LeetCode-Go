package leetcode

func divide(dividend int, divisor int) int {
	ans, sign := 0, 1
	if dividend < 0 {
		dividend, sign = -dividend, -sign
	}
	if divisor < 0 {
		divisor, sign = -divisor, -sign
	}
	// e.g. 1023 / 1 = 512 + 256 + 128 + 64 + 32 + 16 + 8 + 4 + 2 + 1
	for dividend >= divisor {
		// accumulated divisor, times of the original divisor
		acc, mul := divisor, 1
		for acc<<1 <= dividend {
			acc <<= 1
			mul <<= 1
		}
		ans += mul
		dividend -= acc
	}
	if sign == -1 {
		ans = -ans
	}
	if ans > 2147483647 || ans < -2147483648 {
		return 2147483647
	}
	return ans
}
