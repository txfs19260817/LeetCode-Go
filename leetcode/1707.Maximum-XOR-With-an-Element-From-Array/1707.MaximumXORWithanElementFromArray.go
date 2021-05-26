package _707_Maximum_XOR_With_an_Element_From_Array

import "sort"

const L = 30 // nums[i] <= 10^9

func maximizeXor(nums []int, queries [][]int) []int {
	ans, t, j := make([]int, len(queries)), &trie{}, 0
	sort.Ints(nums)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][1] < queries[j][1]
	})
	for _, query := range queries {
		x, m, qIdx := query[0], query[1], query[2]
		for ; j < len(nums) && nums[j] <= m; j++ {
			t.insert(nums[j])
		}
		if j == 0 { // trie is empty
			ans[qIdx] = -1
		} else {
			ans[qIdx] = t.getMaxXor(x)
		}
	}
	return ans
}

type trie struct {
	children [2]*trie
}

func (t *trie) insert(num int) {
	p := t
	for i := L - 1; i >= 0; i-- {
		bit := num >> i & 1
		if p.children[bit] == nil {
			p.children[bit] = &trie{}
		}
		p = p.children[bit]
	}
}

func (t *trie) getMaxXor(num int) (ans int) {
	p := t
	for i := L - 1; i >= 0; i-- {
		bit := num >> i & 1
		if p.children[bit^1] != nil {
			ans |= 1 << i
			bit ^= 1
		}
		p = p.children[bit]
	}
	return
}
