package leetcode

func countCharacters(words []string, chars string) int {
	ans, chars2Count := 0, map[rune]int{}
	for _, r := range chars {
		chars2Count[r]++
	}
	for _, word := range words {
		word2Count, valid := map[rune]int{}, true
		for _, r := range word {
			word2Count[r]++
			if word2Count[r] > chars2Count[r] {
				valid = false
				break
			}
		}
		if valid {
			ans += len(word)
		}
	}
	return ans
}
