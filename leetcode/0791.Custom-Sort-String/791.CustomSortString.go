package leetcode

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

func customSortString2(order string, s string) string {
	bucket := [26]int{}
	for _, c := range s {
		bucket[c-'a']++
	}
	ans := make([]byte, 0, len(s))
	for _, o := range order {
		cnt := bucket[o-'a']
		for range cnt {
			ans = append(ans, byte(o))
		}
		bucket[o-'a'] = 0
	}
	for i := range bucket {
		for bucket[i] > 0 {
			ans = append(ans, byte(i+'a'))
			bucket[i]--
		}
	}
	return string(ans)
}
