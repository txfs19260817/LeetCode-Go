package leetcode

import (
	"strings"
	"sync"
)

type HtmlParser struct {
	graph map[string][]string // url -> 邻接表（它能访问到的 url）
}

func (p *HtmlParser) GetUrls(url string) []string {
	return p.graph[url]
}

/*
This is HtmlParser's API interface.
You should not implement it, or speculate about its implementation

	type HtmlParser struct {
	    maps  map[string]int
	    imaps map[int]string
	    a     map[int][]int
	}
*/
func crawl(startUrl string, htmlParser *HtmlParser) []string {
	hostname := getHostName(startUrl)
	var visited sync.Map // key: string, value: struct{}
	var wg sync.WaitGroup

	var dfs func(string)
	dfs = func(url string) {
		defer wg.Done()
		for _, next := range htmlParser.GetUrls(url) {
			if getHostName(next) != hostname {
				continue
			}
			if _, loaded := visited.LoadOrStore(next, struct{}{}); !loaded {
				wg.Add(1)
				go dfs(next)
			}
		}
	}

	visited.Store(startUrl, struct{}{})
	wg.Add(1)
	go dfs(startUrl)
	wg.Wait()

	var ans []string
	visited.Range(func(key, value interface{}) bool {
		ans = append(ans, key.(string))
		return true
	})
	return ans
}

func getHostName(url string) string {
	hostname := strings.Split(strings.TrimPrefix(url, "http://"), "/")[0]
	return strings.Split(hostname, ":")[0]
}
