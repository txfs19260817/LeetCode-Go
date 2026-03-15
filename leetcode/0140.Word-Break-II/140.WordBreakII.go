package leetcode

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0
	for _, word := range wordDict {
		dict[word] = struct{}{}
		maxLen = max(maxLen, len(word))
	}

	memo := make(map[int][]string, len(s)+1)
	var dfs func(int) []string
	dfs = func(start int) []string {
		// memo[start] caches all sentences that can be formed from s[start:].
		if sentences, ok := memo[start]; ok {
			return sentences
		}
		if start == len(s) {
			// Empty suffix acts as the base case for sentence concatenation.
			return []string{""}
		}

		var sentences []string
		endLimit := start + maxLen
		if endLimit > len(s) {
			endLimit = len(s)
		}
		for end := start + 1; end <= endLimit; end++ {
			word := s[start:end]
			if _, ok := dict[word]; !ok {
				continue
			}
			for _, suffix := range dfs(end) {
				// Avoid a leading space when the current word is the tail.
				if suffix == "" {
					sentences = append(sentences, word)
				} else {
					sentences = append(sentences, word+" "+suffix)
				}
			}
		}

		memo[start] = sentences
		return sentences
	}

	return dfs(0)
}

func wordBreak2(s string, wordDict []string) []string {
	// dp[i] stores all valid sentences that compose s[:i].
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			wLen := len(word)
			if i < wLen || s[i-wLen:i] != word {
				continue
			}
			for _, prev := range dp[i-wLen] {
				// Append the current word to every sentence ending right before it.
				if prev == "" {
					dp[i] = append(dp[i], word)
				} else {
					dp[i] = append(dp[i], prev+" "+word)
				}
			}
		}
	}

	return dp[len(s)]
}
