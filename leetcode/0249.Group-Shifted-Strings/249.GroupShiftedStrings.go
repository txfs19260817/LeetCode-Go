package leetcode

func groupStrings(strings []string) [][]string {
	origin2idx := map[string][]string{}
	for _, s := range strings {
		offset, bytes := int(s[0]-'a'), make([]byte, len(s))
		for j, c := range s {
			fixed := int(c-'a') - offset
			if fixed < 0 {
				fixed += 26
			}
			bytes[j] = byte(fixed)
		}
		origin := string(bytes)
		origin2idx[origin] = append(origin2idx[origin], s)
	}
	var ans [][]string
	for _, sSlice := range origin2idx {
		ans = append(ans, sSlice)
	}
	return ans
}
