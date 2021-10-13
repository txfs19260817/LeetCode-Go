package _893_Groups_of_Special_Equivalent_Strings

import "sort"

func numSpecialEquivGroups(words []string) int {
	g := map[[2]string]bool{}
	for _, word := range words {
		odds, evens := make([]byte, 0, len(word)/2+1), make([]byte, 0, len(word)/2+1)
		for i, r := range word {
			if i%2 == 0 {
				evens = append(evens, byte(r))
			} else {
				odds = append(odds, byte(r))
			}
		}
		sort.Slice(odds, func(i, j int) bool { return odds[i] < odds[j] })
		sort.Slice(evens, func(i, j int) bool { return evens[i] < evens[j] })
		g[[2]string{string(odds), string(evens)}] = true
	}
	return len(g)
}
