package _047_Remove_All_Adjacent_Duplicates_In_String

func removeDuplicates(S string) string {
	var s []rune
	for _, c := range S {
		if len(s) > 0 && s[len(s)-1] == c {
			s = s[:len(s)-1]
		} else {
			s = append(s, c)
		}
	}
	return string(s)
}
