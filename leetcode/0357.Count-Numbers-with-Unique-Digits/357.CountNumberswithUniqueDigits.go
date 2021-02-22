package _357_Count_Numbers_with_Unique_Digits

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	ans, uniNums := 10, 9
	for remain := 9; n > 1 && remain > 0; remain-- {
		uniNums *= remain
		ans += uniNums
		n--
	}
	return ans
}
