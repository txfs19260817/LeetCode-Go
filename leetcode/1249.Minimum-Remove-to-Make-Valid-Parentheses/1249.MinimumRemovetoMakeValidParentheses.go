package _249_Minimum_Remove_to_Make_Valid_Parentheses

func minRemoveToMakeValid(s string) string {
	chars, balance, left := make([]rune, 0, len(s)), 0, 0
	for _, c := range s { // remove invalid ')'
		if c == '(' {
			balance++
			left++
		} else if c == ')' {
			if balance == 0 {
				continue
			}
			balance--
		}
		chars = append(chars, c)
	}
	left -= balance // total left - redundant left = keeping left
	ans := make([]rune, 0, len(chars))
	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' {
			left--
			if left < 0 { // remove rightmost '('
				continue
			}
		}
		ans = append(ans, chars[i])
	}
	return string(ans)
}
