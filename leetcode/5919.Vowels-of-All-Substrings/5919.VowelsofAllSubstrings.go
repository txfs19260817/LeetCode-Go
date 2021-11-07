package _919_Vowels_of_All_Substrings

func countVowels(word string) int64 {
	var ans int64
	for i, ch := range word {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			ans += int64(i+1) * int64(len(word)-i)
		}
	}
	return ans
}
