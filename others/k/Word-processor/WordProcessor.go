package k

import (
	"math"
	"strings"
)

func wrapLines(words []string, maxLen int) (ans []string) {
	for i := 0; i < len(words); {
		var sb strings.Builder
		sb.Grow(maxLen)
		for ; i < len(words); i++ {
			if sb.Len()+1+len(words[i]) <= maxLen { // it still has space for a word + '-'
				sb.WriteString(words[i])
				sb.WriteByte('-')
			} else if sb.Len()+len(words[i]) == maxLen { // exact fit
				sb.WriteString(words[i])
			} else { // overflows
				break
			}
		}
		ans = append(ans, strings.Trim(sb.String(), "-"))
	}
	return
}

func reflowAndJustify(lines []string, maxLen int) (ans []string) {
	var words []string
	for _, line := range lines {
		words = append(words, strings.Fields(line)...)
	}

	for i := 0; i < len(words); {
		var cnt int
		var newLineWords []string
		for i < len(words) && cnt+len(words[i])+1 <= maxLen { // it still has space for a word + '-'
			newLineWords = append(newLineWords, words[i])
			cnt += len(words[i]) + 1
			i++
		}
		if i < len(words) && cnt+len(words[i]) == maxLen { // exact fit
			newLineWords = append(newLineWords, words[i])
			cnt += len(words[i])
			i++
		} else {
			cnt-- // drop the last hyphen
		}

		// only one word (assuming no word is longer than maxLen)
		if len(newLineWords) == 1 {
			ans = append(ans, newLineWords[0])
			continue
		}

		// calculate hyphens
		intervals := len(newLineWords) - 1
		remainingSpaces := maxLen - cnt
		hyphenPerInterval := remainingSpaces/intervals + 1 // already assumes cnt = n words + n-1 hyphens
		additionalHyphens := remainingSpaces % intervals

		var sb strings.Builder
		sb.Grow(maxLen)
		for j, word := range newLineWords {
			sb.WriteString(word)
			if j < len(newLineWords)-1 {
				sb.WriteString(strings.Repeat("-", hyphenPerInterval))
				if additionalHyphens > 0 {
					sb.WriteByte('-')
					additionalHyphens--
				}
			}
		}
		ans = append(ans, sb.String())
	}
	return
}

func balancedWrapLines(text string, maxLen int) (ans []string) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}

	// 1. 找出单词中的最大长度。这是我们迭代的下限。
	// 任何一行的长度都不能小于最长单词的长度。
	minPossibleWidth := 0
	for _, w := range words {
		minPossibleWidth = max(minPossibleWidth, len(w))
	}

	// 2. 核心策略：遍历所有可能的“目标宽度”。
	// 最优解的最长行长度一定在 [minPossibleWidth, maxLen] 之间。
	// 我们对每一个可能的宽度 limit 运行一次标准的“最小化参差度”动态规划。
	minBalanceScore := math.MaxInt64
	for limit := minPossibleWidth; limit <= maxLen; limit++ {
		lines := solveDP(words, limit)

		// 计算当前方案的真实平衡分数，如果找到了更小的分数，更新最优解
		if curScore := calculateBalanceScore(lines); curScore < minBalanceScore {
			minBalanceScore, ans = curScore, lines
		}
	}

	return
}

// solveDP 使用动态规划找到在给定 limit 限制下的最优分行方式。
// 这里优化的目标是：Minimize Sum((limit - lineLength)^2)
func solveDP(words []string, limit int) []string {
	n := len(words)
	// dp[i] 表示从第 i 个单词开始到结束的最小代价
	dp := make([]int, n+1)
	// split[i] 记录从第 i 个单词开始的最佳分行点（即下一行的起始单词索引）
	split := make([]int, n+1)

	// 初始化 DP 数组
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt64
	}
	dp[n] = 0

	// 从后往前计算 DP
	for i := n - 1; i >= 0; i-- {
		currentLen := -1 // 初始化为 -1 是为了抵消第一个单词前多加的一个空格长度

		for j := i; j < n; j++ {
			// 加上当前单词的长度，以及它前面的空格（如果是行首单词，-1+1+len = len）
			currentLen += 1 + len(words[j])
			if currentLen > limit { // 如果超过限制，这一行不能再加词了
				break
			}

			// 如果剩余部分的代价是无穷大，说明后面无法排版，跳过
			if dp[j+1] == math.MaxInt64 {
				continue
			}
			// 计算当前行的代价：(limit - actual_length)^2
			cost := (limit - currentLen) * (limit - currentLen)
			// 总代价更小则更新下一行的起始单词索引
			if totalCost := cost + dp[j+1]; dp[i] > totalCost {
				dp[i] = totalCost
				split[i] = j + 1
			}
		}
	}

	// 重建分行结果
	var lines []string

	for curr := 0; curr < n; {
		next := split[curr]
		lineStr := strings.Join(words[curr:next], "-") // 将当前行的单词用 "-" 连接
		lines = append(lines, lineStr)
		curr = next
	}

	return lines
}

// calculateBalanceScore 根据题目定义计算平衡分数
// 定义：找出最长行长度 Max。Score = Sum((Max - Line_i_Len)^2)
func calculateBalanceScore(lines []string) int {
	// 1. 找到实际的最长行
	maxLineLen := 0
	for _, line := range lines {
		maxLineLen = max(maxLineLen, len(line))
	}

	// 2. 计算平方差之和
	score := 0
	for _, line := range lines {
		diff := maxLineLen - len(line)
		score += diff * diff
	}
	return score
}
