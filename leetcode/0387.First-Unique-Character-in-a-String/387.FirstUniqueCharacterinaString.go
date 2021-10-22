package _387_First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	rune2cnt := map[rune]int{}
	for _, r := range s {
		rune2cnt[r]++
	}
	for i, r := range s {
		if rune2cnt[r] == 1 {
			return i
		}
	}
	return -1
}
