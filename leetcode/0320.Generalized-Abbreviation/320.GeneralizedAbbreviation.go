package leetcode

import (
	"math"
	"strconv"
)

func generateAbbreviations(word string) []string {
	bitmask := int(math.Pow(2, float64(len(word)))) - 1
	ans := make([]string, bitmask+1)
	for i := 0; i <= bitmask; i++ {
		count, w := 0, make([]byte, 0, len(word))
		for j := 0; j < len(word); j++ {
			if b := i >> j; b&1 == 1 {
				count++
			} else if b&1 == 0 {
				if count != 0 {
					w = append(w, []byte(strconv.Itoa(count))...)
					count = 0
				}
				w = append(w, word[j])
			}
		}
		if count != 0 {
			w = append(w, []byte(strconv.Itoa(count))...)
		}
		ans[i] = string(w)
	}
	return ans
}
