package _692_Top_K_Frequent_Words

import "sort"

func topKFrequent(words []string, k int) []string {
	ans, word2freq, freq2word := make([]string, 0, k), map[string]int{}, make([][]string, len(words))
	for _, word := range words {
		word2freq[word]++
	}
	for w, f := range word2freq {
		if idx := sort.SearchStrings(freq2word[f], w); idx == len(freq2word[f]) {
			freq2word[f] = append(freq2word[f], w)
		} else {
			freq2word[f] = append(freq2word[f][:idx+1], freq2word[f][idx:]...)
			freq2word[f][idx] = w
		}
	}
	for i := len(freq2word) - 1; i >= 0 && len(ans) < k; i-- {
		if k-len(ans) > len(freq2word[i]) {
			ans = append(ans, freq2word[i]...)
			continue
		}
		for j := 0; len(ans) < k; j++ {
			ans = append(ans, freq2word[i][j])
		}
	}
	return ans
}
