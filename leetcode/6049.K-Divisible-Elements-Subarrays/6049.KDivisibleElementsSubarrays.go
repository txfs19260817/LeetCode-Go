package _049_K_Divisible_Elements_Subarrays

func countDistinct(nums []int, k int, p int) int {
	ans, prefix, trie := 0, make([]int, len(nums)), NewTrie()
	for i, num := range nums {
		if num%p == 0 {
			prefix[i] = 1
		}
	}
	for i := 1; i < len(prefix); i++ {
		prefix[i] += prefix[i-1]
	}
	prefix = append([]int{0}, prefix...)
	for i := 0; i < len(prefix)-1; i++ {
		for j := i + 1; j < len(prefix); j++ {
			if prefix[j]-prefix[i] <= k {
				if trie.Insert(nums[i:j]) {
					ans++
				}
			}
		}
	}
	return ans
}

type Trie struct {
	children [201]*Trie
	isEnd    bool
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word []int) (isNew bool) {
	p := t
	for _, ch := range word {
		if p.children[ch] == nil {
			p.children[ch] = &Trie{}
		}
		p = p.children[ch]
	}
	if !p.isEnd {
		p.isEnd = true
		isNew = true
	}
	return
}
