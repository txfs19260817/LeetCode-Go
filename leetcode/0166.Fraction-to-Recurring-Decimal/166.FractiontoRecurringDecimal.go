package _166_Fraction_to_Recurring_Decimal

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var ans []byte
	if numerator > 0 && denominator < 0 || numerator < 0 && denominator > 0 {
		ans = append(ans, '-')
	}
	numerator, denominator = abs(numerator), abs(denominator)
	ans = append(ans, []byte(strconv.Itoa(numerator/denominator))...)
	reminder := numerator % denominator
	if reminder != 0 {
		ans = append(ans, '.')
	}
	rem2idx := map[int]int{}
	for reminder != 0 {
		if i, ok := rem2idx[reminder]; ok {
			ans = append(ans[:i+1], ans[i:]...)
			ans[i] = '('
			ans = append(ans, ')')
			break
		}
		rem2idx[reminder] = len(ans)
		reminder *= 10
		ans = append(ans, []byte(strconv.Itoa(reminder/denominator))...)
		reminder %= denominator
	}
	return string(ans)
}

func abs(a int) int {
	var b = int64(a)
	if b < 0 {
		return int(-1 * b)
	}
	return int(b)
}
