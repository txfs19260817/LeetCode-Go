package _129_Capitalize_the_Title

import (
	"strings"
)

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	converted := make([]string, len(words))
	for i, word := range words {
		word = strings.ToLower(word)
		if len(word) > 2 {
			bytes := []byte(word)
			if bytes[0] >= 'a' {
				bytes[0] -= 32
			}
			converted[i] = string(bytes)
		} else {
			converted[i] = word
		}
	}
	return strings.Join(converted, " ")
}
