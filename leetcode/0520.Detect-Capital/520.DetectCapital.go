package _520_Detect_Capital

func detectCapitalUse(word string) bool {
	var upper int
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			upper++
		}
	}
	return (word[0] >= 'A' && word[0] <= 'Z' && upper == 1) || upper == 0 || upper == len(word)
}
