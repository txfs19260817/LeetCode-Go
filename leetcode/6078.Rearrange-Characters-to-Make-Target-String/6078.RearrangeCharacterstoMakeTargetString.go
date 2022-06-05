package _078_Rearrange_Characters_to_Make_Target_String

func rearrangeCharacters(s string, target string) int {
	m, ans := map[rune]int{}, 0
	for _, r := range s {
		m[r]++
	}
	for {
		for _, r := range target {
			if m[r] == 0 {
				return ans
			}
			m[r]--
		}
		ans++
	}
}
