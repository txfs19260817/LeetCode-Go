package _524_Longest_Word_in_Dictionary_through_Deleting

func findLongestWord(s string, d []string) string {
	var ans string
	for _, ds := range d {
		if len(ds) > len(s) || len(ds) < len(ans) || len(ds) == len(ans) && ds >= ans {
			continue
		}
		var j int
		for i := 0; i < len(s) && j < len(ds); i++ {
			if s[i] == ds[j] {
				j++
			}
		}
		if j == len(ds) {
			ans = ds
		}
	}
	return ans
}
