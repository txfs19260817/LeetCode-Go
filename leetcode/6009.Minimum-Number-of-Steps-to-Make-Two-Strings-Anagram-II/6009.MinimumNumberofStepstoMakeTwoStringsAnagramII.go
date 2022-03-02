package _009_Minimum_Number_of_Steps_to_Make_Two_Strings_Anagram_II

func minSteps(s string, t string) int {
	ans, m := 0, map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
	}
	for _, v := range m {
		ans += abs(v)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
