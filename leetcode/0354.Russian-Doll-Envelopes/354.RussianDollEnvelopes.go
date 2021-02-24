package _354_Russian_Doll_Envelopes

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	var seq []int
	for _, e := range envelopes {
		i := sort.Search(len(seq), func(i int) bool {
			return seq[i] >= e[1]
		})
		if i == len(seq) {
			seq = append(seq, e[1])
		} else {
			seq[i] = e[1]
		}
	}
	return len(seq)
}
