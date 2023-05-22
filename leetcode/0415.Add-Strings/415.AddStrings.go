package leetcode

import "strconv"

func addStrings(num1 string, num2 string) string {
	var ans string
	var carry int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
		}
		ans = strconv.Itoa(sum%10) + ans
		carry = sum / 10
	}
	return ans
}
