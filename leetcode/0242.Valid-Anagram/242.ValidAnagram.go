package _242_Valid_Anagram

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
