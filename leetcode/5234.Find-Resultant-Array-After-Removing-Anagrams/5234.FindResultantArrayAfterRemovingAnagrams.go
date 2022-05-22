package _234_Find_Resultant_Array_After_Removing_Anagrams

import "sort"

func removeAnagrams(words []string) []string {
	var ans []string
	for i := 0; i < len(words); {
		sortedWord := sortStr(words[i])
		ans = append(ans, words[i])
		j := i + 1
		for ; j < len(words); j++ {
			wj := sortStr(words[j])
			if sortedWord != wj {
				break
			}
		}
		i = j
	}
	return ans
}

func sortStr(word string) string {
	b := []byte(word)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
