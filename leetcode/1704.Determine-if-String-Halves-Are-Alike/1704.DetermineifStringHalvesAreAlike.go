package _704_Determine_if_String_Halves_Are_Alike

func halvesAreAlike(s string) bool {
	var count int
	for i := 0; i < len(s)/2; i++ {
		if isVowel(s[i]) {
			count++
		}
		if isVowel(s[i+len(s)/2]) {
			count--
		}
	}
	return count == 0
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
