package _928_Decode_the_Slanted_Ciphertext

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	ans, cols := make([]byte, 0, len(encodedText)), len(encodedText)/rows
	rowPtrs := make([]int, rows)
	for i := range rowPtrs {
		rowPtrs[i] = i * (cols + 1)
	}
	for c := 0; c < cols; c++ {
		for i, ptr := range rowPtrs {
			if ptr < len(encodedText) && ptr < (i+1)*(cols+1) {
				ans = append(ans, encodedText[ptr])
				rowPtrs[i]++
			}
		}
	}
	return strings.TrimRight(string(ans), " ")
}
