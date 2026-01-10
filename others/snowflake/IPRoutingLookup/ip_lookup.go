package snowflake

import (
	"fmt"
	"strconv"
	"strings"
)

type Trie struct {
	children [2]*Trie
	target   string
}

func (t *Trie) Insert(ip, target string) {
	p := t
	for _, c := range ip {
		i := c - '0'
		if p.children[i] == nil {
			p.children[i] = &Trie{}
		}
		p = p.children[i]
	}
	p.target = target
}

func (t *Trie) Search(ip string) string {
	p := t
	best := p.target
	for _, c := range ip {
		i := c - '0'
		if p.children[i] == nil {
			return best
		}
		p = p.children[i]
		if len(p.target) > 0 {
			best = p.target
		}
	}
	return best
}

func routeLookup(routes [][]string, queries []string) []string {
	t := &Trie{}
	for _, route := range routes {
		cidr, target := route[0], route[1]
		parts := strings.Split(cidr, "/")
		ip, maskStr := ip2bin(parts[0]), parts[1]
		maskLen, _ := strconv.Atoi(maskStr)
		t.Insert(ip[:maskLen], target)
	}

	ans := make([]string, 0, len(queries))
	for _, q := range queries {
		ans = append(ans, t.Search(ip2bin(q)))
	}
	return ans
}

func ip2bin(ip string) (ans string) {
	var a int
	for _, p := range strings.Split(ip, ".") {
		v, _ := strconv.Atoi(p)
		a = a<<8 | v
	}
	return fmt.Sprintf("%032b", a)
}
