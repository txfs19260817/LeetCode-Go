package leetcode

func prefixCount(words []string, pref string) int {
	var ans int
	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}
		var flag bool
		for i := range pref {
			if pref[i] != word[i] {
				flag = true
				break
			}
		}
		if !flag {
			ans++
		}
	}
	return ans
}
