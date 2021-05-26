package _244_Shortest_Word_Distance_II

import "sort"

type WordDistance struct {
	word2indices map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	word2indices := make(map[string][]int)
	for i, s := range wordsDict {
		word2indices[s] = append(word2indices[s], i)
	}
	return WordDistance{word2indices}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	l1, l2 := this.word2indices[word1], this.word2indices[word2]
	ans := abs(l1[0] - l2[0])
	for _, i := range l1 {
		switch j := sort.SearchInts(l2, i); j {
		case 0:
			ans = min(ans, abs(i-l2[0]))
		case len(l2):
			ans = min(ans, abs(i-l2[len(l2)-1]))
		default:
			ans = min(ans, min(abs(l2[j]-i), abs(l2[j-1]-i)))
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
