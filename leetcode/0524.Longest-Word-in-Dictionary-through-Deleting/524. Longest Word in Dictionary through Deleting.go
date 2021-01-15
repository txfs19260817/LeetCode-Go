package _524_Longest_Word_in_Dictionary_through_Deleting

func findLongestWord(s string, d []string) string {
	if len(s) == 0 || len(d) == 0 {
		return ""
	}
	res := ""
	for _, ds := range d {
		if compare(s, ds) && (len(ds) > len(res) || len(ds) == len(res) && res > ds) {
			res = ds
		}
	}

	return res
}

func compare(s, ds string) bool {
	if len(ds) > len(s) || len(ds) == 0 {
		return false
	}

	j := 0
	for i := 0; i < len(s) && j < len(ds); i++ {
		if s[i] == ds[j] {
			j++
		}
	}
	if j == len(ds) {
		return true
	}
	return false
}
