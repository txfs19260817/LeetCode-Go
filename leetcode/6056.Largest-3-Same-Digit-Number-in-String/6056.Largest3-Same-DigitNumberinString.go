package leetcode

import (
	"fmt"
)

func largestGoodInteger(num string) string {
	for i := 9; i >= 0; i-- {
		sub := fmt.Sprintf("%d%d%d", i, i, i)
		if num[len(num)-3:] == sub {
			return sub
		}
		for j := 0; j < len(num)-3; j++ {
			if sub == num[j:j+3] && (j+3 == len(num) || byte(i+'0') != num[j+3]) {
				return sub
			}
		}
	}
	return ""
}
