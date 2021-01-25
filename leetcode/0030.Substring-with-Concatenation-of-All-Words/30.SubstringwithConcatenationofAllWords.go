package _030_Substring_with_Concatenation_of_All_Words

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/golang-hua-dong-chuang-kou-fa-by-bloodborne-3/
func findSubstring(s string, words []string) []int {
	wordLen, totalLen := len(words[0]), len(words[0])*len(words)
	var res []int
	if len(s) < totalLen {
		return res
	}
	dict := map[string]int{}
	for _, word := range words {
		dict[word]++
	}
	for start := 0; start < wordLen; start++ {
		l, r, match := start, start, 0
		window := map[string]int{}
		for r+wordLen <= len(s) {
			inWord := s[r : r+wordLen]
			r += wordLen
			if dict[inWord] > 0 {
				window[inWord]++
				if dict[inWord] == window[inWord] {
					match++
				}
			}
			if r-l == totalLen {
				if match == len(dict) {
					res = append(res, l)
				}
				outWord := s[l : l+wordLen]
				l += wordLen
				if dict[outWord] > 0 {
					if dict[outWord] == window[outWord] {
						match--
					}
					window[outWord]--
				}
			}
		}
	}
	return res
}

func findSubstring1(s string, words []string) []int {
	wordLen, totalLen := len(words[0]), len(words[0])*len(words)
	var res []int
	if len(s) < totalLen {
		return res
	}
	dict := map[string]int{}
	for _, word := range words {
		dict[word]++
	}
	tempDict := copyMap(dict)
	for i, start := 0, 0; i < len(s)-wordLen+1 && start < len(s)-wordLen+1; {
		subStrWord := s[i : i+wordLen]
		if tempDict[subStrWord] > 0 {
			tempDict[subStrWord]--
			if isEmptyMap(tempDict) && i+wordLen-start == totalLen {
				res = append(res, start)
				i++
			} else {
				i += wordLen
			}
		} else {
			start++
			i = start
			tempDict = copyMap(dict)
		}
	}
	return res
}

func isEmptyMap(m map[string]int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func copyMap(m map[string]int) map[string]int {
	res := map[string]int{}
	for k, v := range m {
		res[k] = v
	}
	return res
}
