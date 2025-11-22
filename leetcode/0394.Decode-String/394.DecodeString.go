package leetcode

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	numIdxStart, numIdxEnd, backBracketIdx, bracketDepth := -1, -1, 0, 0
	for i, c := range s {
		if c >= '0' && c <= '9' {
			numIdxStart, numIdxEnd = i, i
			for j := i + 1; j < len(s); j++ {
				if s[j] >= '0' && s[j] <= '9' {
					numIdxEnd = j
				} else {
					break
				}
			}
			break
		}
	}
	if numIdxStart == -1 {
		return s
	}
	for i := numIdxEnd + 1; i < len(s); i++ {
		if s[i] == '[' {
			bracketDepth++
		} else if s[i] == ']' {
			bracketDepth--
			if bracketDepth == 0 {
				backBracketIdx = i
				break
			}
		}
	}
	ans := decodeString(s[:numIdxStart])
	mid := decodeString(s[numIdxEnd+2 : backBracketIdx])
	t, _ := strconv.Atoi(s[numIdxStart : numIdxEnd+1])
	for i := 0; i < t; i++ {
		ans += mid
	}
	ans += decodeString(s[backBracketIdx+1:])
	return ans
}

func decodeString2(s string) string {
	ans, num := "", 0
	for i := 0; i < len(s); {
		// chars
		for ; i < len(s) && s[i] >= 'a' && s[i] <= 'z'; i++ {
			ans += string(s[i])
		}
		// num
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]-'0')
		}
		if i == len(s) {
			break
		}
		i++
		j := i
		for depth := 1; depth > 0; j++ {
			if s[j] == '[' {
				depth++
			} else if s[j] == ']' {
				depth--
			}
		}
		ans += strings.Repeat(decodeString(s[i:j-1]), num)
		i = j
		num = 0
	}
	return ans
}
