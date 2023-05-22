package leetcode

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	ans := make([]byte, 0, len(encodedText))
	for i, j, k, cols := 0, 0, 0, len(encodedText)/rows; j < cols; {
		ans = append(ans, encodedText[i*cols+j])
		i++
		j++
		if i == rows || j == cols { // reached edge, move to the next diag
			k++
			j = k
			i = 0
		}
	}
	return strings.TrimRight(string(ans), " ")
}
