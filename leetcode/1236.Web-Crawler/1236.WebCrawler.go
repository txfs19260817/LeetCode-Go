package leetcode

import "strings"

type HtmlParser interface {
	GetUrls(url string) []string
}

func crawl(startUrl string, htmlParser HtmlParser) []string {
	visited := map[string]bool{startUrl: true}
	hostname := getHostName(startUrl)
	q := []string{startUrl}
	for len(q) > 0 {
		nextUrl := q[0]
		q = q[1:]
		for _, s := range htmlParser.GetUrls(nextUrl) {
			if getHostName(s) == hostname && !visited[s] {
				visited[s] = true
				q = append(q, s)
			}
		}
	}
	ans := make([]string, 0, len(visited))
	for k := range visited {
		ans = append(ans, k)
	}
	return ans
}

func getHostName(url string) string {
	hostname := strings.Split(strings.TrimPrefix(url, "http://"), "/")[0]
	return strings.Split(hostname, ":")[0]
}
