package _967_Check_if_All_As_Appears_Before_All_Bs

func checkString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 'b' && s[i+1] == 'a' {
			return false
		}
	}
	return true
}
