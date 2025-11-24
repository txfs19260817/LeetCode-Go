package k

import (
	"sort"
	"strconv"
)

func grouping(events [][]string, n int) [][]string {
	// adj[u] 存储 u 连接的所有邻居
	// 这样 len(adj[u]) 就是 u 的当前度数
	adj := map[string]map[string]bool{}

	for _, e := range events {
		action, u, v := e[0], e[1], e[2]

		// 确保 u, v 都在 map 中初始化，即使它们目前没有连接（可能被断开过）
		if _, ok := adj[u]; !ok {
			adj[u] = make(map[string]bool)
		}
		if _, ok := adj[v]; !ok {
			adj[v] = make(map[string]bool)
		}

		if action == "CONNECT" {
			adj[u][v] = true
			adj[v][u] = true
		} else if action == "DISCONNECT" {
			delete(adj[u], v)
			delete(adj[v], u)
		}
	}

	lessThanN, nOrMore := []string{}, []string{}

	// 直接遍历 adj map
	for user, neighbors := range adj {
		// len(neighbors) 即为连接数
		if len(neighbors) < n {
			lessThanN = append(lessThanN, user)
		} else {
			nOrMore = append(nOrMore, user)
		}
	}

	sort.Strings(lessThanN)
	sort.Strings(nOrMore)

	return [][]string{lessThanN, nOrMore}
}

func recommendations(user string, ratings [][]string) []string {
	// 1. 构建数据结构 watched[UserName][MovieName] = Rating
	watched := make(map[string]map[string]int)

	// 记录每个人看过的电影集合，方便快速判断"是否看过"
	for _, r := range ratings {
		u, m, s := r[0], r[1], r[2]
		score, _ := strconv.Atoi(s)

		if _, ok := watched[u]; !ok { // no nil map
			watched[u] = make(map[string]int)
		}
		watched[u][m] = score
	}

	userWatched := watched[user]

	// 推荐列表（用 map 去重）
	recommendSet := make(map[string]bool)

	// 2. 遍历其他用户寻找相似用户
	for otherUser, otherWatched := range watched {
		if otherUser == user {
			continue
		}

		// 判断相似性：是否有共同好评电影 (>3)
		isSimilar := false
		for m, score := range otherWatched {
			// 检查目标用户是否也看过这部电影
			if targetScore, ok := userWatched[m]; ok {
				// 两人都看过，且都给了好评
				if score > 3 && targetScore > 3 {
					isSimilar = true
					break // 只要找到一个就足够证明相似
				}
			}
		}

		// 3. 如果相似，收集推荐
		if isSimilar {
			for m := range otherWatched {
				// 推荐条件：对方看过，但目标用户没看过
				if _, seen := userWatched[m]; !seen {
					recommendSet[m] = true
				}
			}
		}
	}

	// 4. 格式化输出
	ans := make([]string, 0, len(recommendSet))
	for m := range recommendSet {
		ans = append(ans, m)
	}
	return ans
}
