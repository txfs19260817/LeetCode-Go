package _221_Split_a_String_in_Balanced_Strings

func balancedStringSplit(s string) int {
	var ans int
	for i, cnt := 0, 0; i < len(s); i++ {
		if s[i] == 'R' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ans++
		}
	}
	return ans
}
