package leetcode

func characterReplacement(s string, k int) int {
	l, ans, freq, maxFreq := 0, 0, map[byte]int{}, 0
	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		maxFreq = max(maxFreq, freq[s[r]])
		if r+1-l-maxFreq > k {
			freq[s[l]]--
			l++
		}
		ans = max(ans, r+1-l)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement1(s string, k int) int {
	if len(s) < 2 || len(s) == k {
		return len(s)
	}
	ans, freq := 0, map[byte]int{}
	for l, r := 0, 0; l < len(s); {
		if r < len(s) && r-l-getMaxValue(freq, s[r]) < k { // r+1-l-getMaxValue(freq, s[r]) <= k
			freq[s[r]]++
			r++
		} else {
			freq[s[l]]--
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	return ans
}

func getMaxValue(m map[byte]int, extra byte) int {
	var ans int
	var key byte
	for k, v := range m {
		if ans < v {
			ans, key = v, k
		}
	}
	if ans == 0 {
		return 1
	}
	if key == extra {
		ans++
	}
	return ans
}
