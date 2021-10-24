package _906_Number_of_Valid_Words_in_a_Sentence

import (
	"regexp"
	"strings"
)

func countValidWords(sentence string) int {
	var ans int
	lowerAlphabets := regexp.MustCompile(`^[a-z]*([a-z]-[a-z])?[a-z]*[!.,]?$`)
	for _, s := range strings.Split(sentence, " ") {
		if word := strings.TrimSpace(s); len(word) > 0 &&lowerAlphabets.MatchString(word) {
			ans++
		}
	}
	return ans
}
