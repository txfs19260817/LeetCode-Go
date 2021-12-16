package _273_Integer_to_English_Words

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	singles := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands := []string{"", "Thousand", "Million", "Billion"}
	sb := strings.Builder{}
	cvt := func(n int) {
		if n >= 100 {
			sb.WriteString(singles[n/100])
			sb.WriteString(" Hundred ")
			n %= 100
		}
		if n >= 20 {
			sb.WriteString(tens[n/10])
			sb.WriteByte(' ')
			n %= 10
		}
		if n >= 10 {
			sb.WriteString(teens[n-10])
			sb.WriteByte(' ')
		} else if n > 0 && n < 10 {
			sb.WriteString(singles[n])
			sb.WriteByte(' ')
		}
	}
	for i, unit := 3, int(1e9); i >= 0; i, unit = i-1, unit/1000 {
		if num < unit {
			continue
		}
		curNum := num / unit
		num -= curNum * unit
		cvt(curNum)
		sb.WriteString(thousands[i])
		sb.WriteByte(' ')
	}
	return strings.TrimSpace(sb.String())
}
