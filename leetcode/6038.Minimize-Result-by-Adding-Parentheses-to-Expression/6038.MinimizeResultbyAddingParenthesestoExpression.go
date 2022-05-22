package _038_Minimize_Result_by_Adding_Parentheses_to_Expression

import (
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	ans := "(" + expression + ")"
	plusIdx := strings.Index(expression, "+")
	as, bs := expression[:plusIdx], expression[plusIdx+1:]
	anum := atoi(as, 0)
	bnum := atoi(bs, 0)
	mini := anum + bnum
	for l := 0; l < len(as); l++ {
		for r := 1; r <= len(bs); r++ {
			v := atoi(as[:l], 1) * (atoi(as[l:], 0) + atoi(bs[:r], 0)) * atoi(bs[r:], 1)
			if v < mini {
				mini = v
				ans = as[:l] + "(" + as[l:] + "+" + bs[:r] + ")" + bs[r:]
			}
		}
	}
	return ans
}

func atoi(s string, def int) (v int) {
	if len(s) == 0 {
		return def
	}
	v, _ = strconv.Atoi(s)
	return v
}
