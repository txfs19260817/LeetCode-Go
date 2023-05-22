package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ans, chars, line, dir := "", make([][]rune, numRows), 0, 1
	for _, c := range s {
		if line == len(chars)-1 {
			dir = -1
		} else if line == 0 {
			dir = 1
		}
		chars[line] = append(chars[line], c)
		line += dir
	}
	for _, char := range chars {
		ans += string(char)
	}
	return ans
}
