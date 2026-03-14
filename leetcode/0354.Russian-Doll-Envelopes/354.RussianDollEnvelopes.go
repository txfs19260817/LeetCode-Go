package leetcode

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	var seq []int
	for _, e := range envelopes {
		if i := sort.SearchInts(seq, e[1]); i == len(seq) {
			seq = append(seq, e[1])
		} else {
			seq[i] = e[1]
		}
	}
	return len(seq)
}
