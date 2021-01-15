package _767_Reorganize_String

import (
	"sort"
)

func reorganizeString(S string) string {
	m := map[rune]int{}
	for _, c := range S {
		m[c]++
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	chars := make([]rune, 0, len(S))
	for _, kv := range ss {
		for i := 0; i < kv.Value; i++ {
			chars = append(chars, kv.Key)
		}
	}

	res := make([]rune, len(chars))
	idx := -2
	for _, char := range chars {
		idx += 2
		if idx >= len(chars) {
			idx = 1
		}
		if (idx-1 >= 0 && char == res[idx-1]) || (idx+1 <= len(res)-1 && char == res[idx+1]) {
			return ""
		}
		res[idx] = char
	}
	return string(res)
}

type kv struct {
	Key   rune
	Value int
}
