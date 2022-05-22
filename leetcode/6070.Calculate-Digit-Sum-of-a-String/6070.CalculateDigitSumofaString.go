package _070_Calculate_Digit_Sum_of_a_String

import "strconv"

func digitSum(s string, k int) string {
	bs := []byte(s)
	for len(bs) > k {
		var q []byte
		for i := 0; i < len(bs); i += k {
			var grp int
			for j := i; j < i+k && j < len(bs); j++ {
				grp += int(bs[j] - '0')
			}
			q = append(q, strconv.Itoa(grp)...)
		}
		bs = q
	}
	return string(bs)
}
