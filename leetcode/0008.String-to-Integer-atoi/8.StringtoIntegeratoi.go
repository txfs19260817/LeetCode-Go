package _008_String_to_Integer_atoi

func myAtoi(s string) int {
	ans, begin, sym := 0, 0, 1
	for ; begin < len(s) && s[begin] == ' '; begin++ {
	}
	if begin == len(s) {
		return 0
	}
	if !(s[begin] <= '9' && s[begin] >= '0' || s[begin] == '+' || s[begin] == '-') {
		return 0
	}
	if s[begin] == '+' || s[begin] == '-' {
		if s[begin] == '-' {
			sym = -1
		}
		begin++
	}
	for ; begin < len(s) && s[begin] <= '9' && s[begin] >= '0'; begin++ {
		digit := int(s[begin] - '0')
		if sym == 1 && ans > (2147483647-digit)/10 {
			return 2147483647
		}
		if sym == -1 && -ans < (-2147483648+digit)/10 {
			return -2147483648
		}
		ans *= 10
		ans += digit
	}
	if sym == -1 {
		ans = -ans
	}
	return ans
}
