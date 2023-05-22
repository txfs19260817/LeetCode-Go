package leetcode

func minAddToMakeValid(s string) int {
	var balance, correct int
	for _, c := range s {
		if c == '(' {
			balance++
		} else if balance > 0 {
			balance--
			correct += 2
		}
	}
	return len(s) - correct
}
