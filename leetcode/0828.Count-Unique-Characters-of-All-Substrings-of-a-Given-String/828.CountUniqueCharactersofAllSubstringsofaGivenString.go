package _828_Count_Unique_Characters_of_All_Substrings_of_a_Given_String

func uniqueLetterString(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		l := i - 1
		for l >= 0 && s[i] != s[l] {
			l--
		}
		r := i + 1
		for r < len(s) && s[i] != s[r] {
			r++
		}
		res += (i - l) * (r - i)
	}
	return res
}
