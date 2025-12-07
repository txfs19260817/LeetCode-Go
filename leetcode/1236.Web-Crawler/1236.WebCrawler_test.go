package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHtmlParser struct {
	urlsEdges map[string][]string
}

func (m *MockHtmlParser) GetUrls(url string) []string {
	return m.urlsEdges[url]
}

func Test_crawl(t *testing.T) {
	tests := []struct {
		name     string
		urls     []string
		edges    [][]int
		startUrl string
		want     []string
	}{
		{
			name: "Example 1",
			urls: []string{
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
				"http://news.google.com",
				"http://news.yahoo.com/us",
			},
			edges:    [][]int{{2, 0}, {2, 1}, {3, 2}, {3, 1}, {0, 4}},
			startUrl: "http://news.yahoo.com/news/topics/",
			want: []string{
				"http://news.yahoo.com/news/topics/",
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/us",
			},
		},
		{
			name: "Example 2",
			urls: []string{
				"http://news.yahoo.com",
				"http://news.yahoo.com/news",
				"http://news.yahoo.com/news/topics/",
				"http://news.google.com",
			},
			edges:    [][]int{{0, 2}, {2, 1}, {3, 2}, {3, 1}, {3, 0}},
			startUrl: "http://news.google.com",
			want:     []string{"http://news.google.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockParser := &MockHtmlParser{urlsEdges: make(map[string][]string)}
			for _, edge := range tt.edges {
				from := tt.urls[edge[0]]
				to := tt.urls[edge[1]]
				mockParser.urlsEdges[from] = append(mockParser.urlsEdges[from], to)
			}

			got := crawl(tt.startUrl, mockParser)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
