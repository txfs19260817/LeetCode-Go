package _043_Multiply_Strings

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var ans string
	ints := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			ints[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	for i := len(ints) - 1; i > 0; i-- {
		ints[i-1] += ints[i] / 10
		ints[i] %= 10
	}
	for _, n := range ints {
		ans += strconv.Itoa(n)
	}
	// trim leading 0
	idx := 0
	for i, c := range ans {
		if c != '0' {
			idx = i
			break
		}
	}
	return ans[idx:]
}
