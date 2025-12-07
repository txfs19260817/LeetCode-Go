package leetcode

import (
	"sort"
	"strings"
)

type Trie struct {
	name     string
	children map[string]*Trie // name to node
	content  strings.Builder
	isFile   bool
}

func (t *Trie) find(path string) *Trie {
	if path == "/" {
		return t
	}
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	p := t
	for _, name := range parts {
		if _, ok := p.children[name]; ok {
			p = p.children[name]
		} else {
			return nil
		}
	}
	return p
}

func (t *Trie) insert(path string, isFile bool) *Trie {
	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")
	p := t
	for i, name := range parts {
		if _, ok := p.children[name]; !ok {
			p.children[name] = &Trie{name: name, children: make(map[string]*Trie), isFile: i == len(parts)-1 && isFile}
		}
		p = p.children[name]
	}
	return p
}

type FileSystem struct {
	root *Trie
}

func Constructor() FileSystem {
	return FileSystem{root: &Trie{name: "/", children: make(map[string]*Trie)}}
}

func (fs *FileSystem) Ls(path string) []string {
	node := fs.root.find(path)
	if node == nil {
		return nil
	}
	if node.isFile {
		return []string{node.name}
	}
	ans := make([]string, 0, len(node.children))
	for _, t := range node.children {
		ans = append(ans, t.name)
	}
	sort.Strings(ans)
	return ans
}

func (fs *FileSystem) Mkdir(path string) {
	fs.root.insert(path, false)
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	t := fs.root.insert(filePath, true)
	t.content.WriteString(content)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	t := fs.root.find(filePath)
	if t != nil && t.isFile {
		return t.content.String()
	}
	return ""
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
