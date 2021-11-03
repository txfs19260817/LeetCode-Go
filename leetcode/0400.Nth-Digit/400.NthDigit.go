package _400_Nth_Digit

import (
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	// calculate how many digits the number has
	digits := 1
	for base := 9; n-base*digits > 0; base *= 10 {
		n -= base * digits
		digits++
	}

	// calculate what the number is, as well as the target index
	number, unit := int(math.Pow10(digits-1))+(n-1)/digits, (n-1)%digits

	// get that digit
	return int([]byte(strconv.Itoa(number))[unit] - '0')
}
