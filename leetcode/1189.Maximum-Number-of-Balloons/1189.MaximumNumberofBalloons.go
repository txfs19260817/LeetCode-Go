package _189_Maximum_Number_of_Balloons

func maxNumberOfBalloons(text string) int {
	m := map[rune]int{}
	for _, c := range text {
		if c == 'b' || c == 'a' || c == 'l' || c == 'o' || c == 'n' {
			m[c]++
		}
	}
	m['l'] /= 2
	m['o'] /= 2
	return min(m['b'], min(m['a'], min(m['l'], min(m['o'], m['n']))))
}

func maxNumberOfBalloons2(text string) int {
	return findMaxNumberofPattern(text, "balloon")
}

func findMaxNumberofPattern(text, pattern string) int {
	ans, mt, mp := 1<<30, map[rune]int{}, map[rune]int{}
	for _, c := range text {
		mt[c]++
	}
	for _, c := range pattern {
		mp[c]++
	}
	for c, n := range mp {
		ans = min(ans, mt[c]/n)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
