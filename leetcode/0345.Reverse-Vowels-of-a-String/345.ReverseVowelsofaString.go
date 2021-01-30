package _345_Reverse_Vowels_of_a_String

func reverseVowels(s string) string {
	chars := []byte(s)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !vowels[chars[l]] {
			l++
		}
		for l < r && !vowels[chars[r]] {
			r--
		}
		chars[l], chars[r] = chars[r], chars[l]
	}
	return string(chars)
}
