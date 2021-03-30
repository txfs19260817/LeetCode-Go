package _049_Group_Anagrams

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	m := map[[26]int][]string{}
	for _, s := range strs {
		var key [26]int
		for _, c := range s {
			key[c-'a']++
		}
		m[key] = append(m[key], s)
	}
	for _, strings := range m {
		ans = append(ans, strings)
	}
	return ans
}
