package sn

import "sort"

func substringCalculator(s string) int {
	suffixArray := buildSuffixArray(s)
	ans := len(suffixArray[0])
	lcpArray := buildLcpArray(suffixArray)
	for i := 1; i < len(suffixArray); i++ {
		ans += len(suffixArray[i]) - lcpArray[i-1]
	}
	return ans
}

func buildSuffixArray(s string) []string {
	suffixArray := make([]string, len(s))
	for i := range s {
		suffixArray[i] = s[i:]
	}
	sort.Strings(suffixArray)
	return suffixArray
}

func buildLcpArray(suffixArray []string) []int {
	commonPrefix := func(a, b string) (l int) {
		n := min(len(a), len(b))
		for i := 0; i < n; i++ {
			if a[i] != b[i] {
				break
			}
			l++
		}
		return
	}
	ans := make([]int, len(suffixArray)-1)
	for i := 0; i < len(suffixArray)-1; i++ {
		ans[i] = commonPrefix(suffixArray[i], suffixArray[i+1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
