package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildParser(urls []string, edges [][]int) *HtmlParser {
	graph := make(map[string][]string)
	for _, e := range edges {
		from := urls[e[0]]
		to := urls[e[1]]
		graph[from] = append(graph[from], to)
	}
	return &HtmlParser{graph: graph}
}

func TestCrawlExample1(t *testing.T) {
	urls := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
		"http://news.yahoo.com/us",
	}
	edges := [][]int{
		{2, 0},
		{2, 1},
		{3, 2},
		{3, 1},
		{0, 4},
	}
	startUrl := "http://news.yahoo.com/news/topics/"

	parser := buildParser(urls, edges)
	got := crawl(startUrl, parser)

	want := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.yahoo.com/us",
	}

	assert.ElementsMatch(t, want, got)
}

func TestCrawlExample2(t *testing.T) {
	urls := []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
	}
	edges := [][]int{
		{0, 2},
		{2, 1},
		{3, 2},
		{3, 1},
		{3, 0},
	}
	startUrl := "http://news.google.com"

	parser := buildParser(urls, edges)
	got := crawl(startUrl, parser)

	want := []string{
		"http://news.google.com",
	}

	assert.ElementsMatch(t, want, got)
}

func TestCrawlLarge(t *testing.T) {
	const N = 50000

	graph := make(map[string][]string)

	// same host urls
	yahooURL := func(i int) string {
		return fmt.Sprintf("http://news.yahoo.com/page/%d", i)
	}

	// cross host url（不应该被 crawl 结果包含）
	otherURL := func(i int) string {
		return fmt.Sprintf("http://other.com/page/%d", i)
	}

	// 构建图
	for i := 0; i < N; i++ {
		from := yahooURL(i)

		// 连向 i+1
		if i+1 < N {
			graph[from] = append(graph[from], yahooURL(i+1))
		}
		// 连向 i+2
		if i+2 < N {
			graph[from] = append(graph[from], yahooURL(i+2))
		}
		// 连向 i/2（制造大量回边/环）
		if i > 0 {
			graph[from] = append(graph[from], yahooURL(i/2))
		}
		// 每个节点再连一条跨域链接，确保不会被爬
		graph[from] = append(graph[from], otherURL(i))
	}

	parser := &HtmlParser{graph: graph}
	startUrl := yahooURL(0)

	got := crawl(startUrl, parser)

	// 1. 长度必须是 N（所有 yahoo host 的节点都应该被访问到）
	assert.Equal(t, N, len(got))

	// 2. 校验结果里没有 cross host
	for _, u := range got {
		assert.Equal(t, "news.yahoo.com", getHostName(u))
	}

	// 3. 校验所有 yahooURL(i) 都在结果里
	sort.Strings(got)
	exists := func(target string) bool {
		idx := sort.SearchStrings(got, target)
		return idx < len(got) && got[idx] == target
	}

	for i := range N {
		assert.True(t, exists(yahooURL(i)))
	}
}
