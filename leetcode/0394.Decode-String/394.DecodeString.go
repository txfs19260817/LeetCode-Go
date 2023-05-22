package leetcode

import "strconv"

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
