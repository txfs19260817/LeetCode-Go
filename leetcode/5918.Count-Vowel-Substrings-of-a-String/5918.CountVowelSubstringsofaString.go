package _918_Count_Vowel_Substrings_of_a_String

func countVowelSubstrings(word string) int {
	var ans int
	isNotVowel := func(b byte) bool {
		return b != 'a' && b != 'e' && b != 'i' && b != 'o' && b != 'u'
	}
	for i := 0; i < len(word); i++ { // split
		if isNotVowel(word[i]) {
			continue
		}
		j := i
		for ; j < len(word); j++ {
			if isNotVowel(word[j]) {
				break
			}
		}
		ans += countValidSub(word[i:j])
		i = j
	}
	return ans
}

func countValidSub(word string) (ans int) {
	if len(word) < 5 {
		return 0
	}
	vowel2cnt := map[byte]int{}
	for l, r := 0, 0; r < len(word)&&l < len(word); {
		vowel2cnt[word[r]]++
		if vowel2cnt['a'] > 0 && vowel2cnt['e'] > 0 && vowel2cnt['i'] > 0 && vowel2cnt['o'] > 0 && vowel2cnt['u'] > 0 {
			ans += len(word) - r
			vowel2cnt[word[l]]--
			l++
			vowel2cnt[word[r]]-- // undo
		} else {
			r++
		}
	}
	return
}
