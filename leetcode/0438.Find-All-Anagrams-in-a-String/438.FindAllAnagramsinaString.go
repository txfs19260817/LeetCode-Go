package _438_Find_All_Anagrams_in_a_String

func findAnagrams(s string, p string) []int {
	var ans []int
	if len(s) < len(p) {
		return ans
	}
	dict, window := map[byte]int{}, map[byte]int{}
	for _, c := range p {
		dict[byte(c)]++
	}
	for l, r := 0, 0; l < len(s)-len(p)+1; {
		if r < len(s) && (r-l < len(p) || !compareMap(dict, window)) {
			window[s[r]]++
			r++
		} else {
			window[s[l]]--
			l++
		}
		if r-l == len(p) && compareMap(dict, window) {
			ans = append(ans, l)
		}
	}
	return ans
}

func compareMap(dict, window map[byte]int) bool {
	for k := range dict {
		if dict[k] > window[k] {
			return false
		}
	}
	return true
}
