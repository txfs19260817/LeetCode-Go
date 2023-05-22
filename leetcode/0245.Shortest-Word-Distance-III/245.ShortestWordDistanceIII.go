package leetcode

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	ans, p1, p2 := 1<<31, -1, -1
	for i, w := range wordsDict {
		if w == word1 {
			if word1 == word2 && p1 != -1 && i-p1 < ans {
				ans = i - p1
			}
			p1 = i
		} else if w == word2 {
			p2 = i
		}
		if p1 != -1 && p2 != -1 {
			if n := abs(p1 - p2); n < ans {
				ans = n
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
