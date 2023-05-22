package leetcode

func countSegments(s string) int {
	var ans int
	for i, c := range s {
		if (i == 0 || s[i-1] == ' ') && c != ' ' {
			ans++
		}
	}
	return ans
}
