package _791_Custom_Sort_String

import "bytes"

func customSortString(order string, s string) string {
	var ans []byte
	var sFreq [26]int
	for _, c := range s {
		sFreq[c-'a']++
	}
	for _, c := range order {
		ans = append(ans, bytes.Repeat([]byte{byte(c)}, sFreq[c-'a'])...)
		sFreq[c-'a'] = 0 // remove chars in `order`
	}
	for _, c := range s {
		if sFreq[c-'a'] != 0 { // append remaining chars
			ans = append(ans, byte(c))
		}
	}
	return string(ans)
}
