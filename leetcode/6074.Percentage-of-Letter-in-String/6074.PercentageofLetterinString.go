package leetcode

func percentageLetter(s string, letter byte) int {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}
	return int(float64(m[rune(letter)]) * 100 / float64(len(s)))
}
