package leetcode

func validWordAbbreviation(word string, abbr string) bool {
	var i, j int
	for i < len(word) && j < len(abbr) {
		if word[i] == abbr[j] {
			i++
			j++
			continue
		}
		if j < len(abbr) && (abbr[j] > '9' || abbr[j] <= '0') {
			return false
		}
		var cnt int
		for j < len(abbr) && abbr[j] <= '9' && abbr[j] >= '0' {
			cnt = cnt*10 + int(abbr[j]-'0')
			j++
		}
		i += cnt
	}
	return i == len(word) && j == len(abbr)
}
