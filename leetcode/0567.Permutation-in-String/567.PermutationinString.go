package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	freq := map[byte]int{}
	for _, c := range s1 {
		freq[byte(c)]++
	}

	for l, r, diff := 0, 0, len(s1); r < len(s2); {
		if freq[s2[r]] > 0 {
			diff--
		}
		freq[s2[r]]--
		r++
		if diff == 0 {
			return true
		}
		if r-l == len(s1) {
			if freq[s2[l]] >= 0 { // s2[l] is in s1
				diff++
			}
			freq[s2[l]]++
			l++
		}
	}
	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	if len(s2) == 0 {
		return false
	}
	dict, window := map[byte]int{}, map[byte]int{}
	for _, c := range s1 {
		dict[byte(c)]++
	}
	for l, r := 0, 0; l < len(s2)-len(s1)+1; {
		if r < len(s2) && (r-l < len(s1) || !compareMap(dict, window)) {
			window[s2[r]]++
			r++
		} else {
			window[s2[l]]--
			l++
		}
		if r-l == len(s1) && compareMap(dict, window) {
			return true
		}
	}
	return false
}

func compareMap(dict, window map[byte]int) bool {
	for k := range dict {
		if dict[k] > window[k] {
			return false
		}
	}
	return true
}
