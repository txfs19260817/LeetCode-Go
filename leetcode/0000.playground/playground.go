package _000_playground

import "sort"

/*
A playground to code benefiting from auto-completion within IDE.
*/
func reorganizeString(S string) string {
	if len(S) < 2 {
		return S
	}
	S = frequencySort(S)
	arr := make([]byte, len(S))
	i := 0
	for j := 0; j < len(S); j++ {
		arr[i] = S[j]
		if i-1 >= 0 && arr[i-1] == arr[i] || i+1 < len(S) && arr[i+1] == arr[i] {
			return ""
		}
		i += 2
		if i >= len(arr) {
			i = 1
		}
	}
	return string(arr)
}

func frequencySort(s string) string {
	if len(s) < 2 {
		return s
	}
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	kvSlice := make([]kv, 0, len(m))
	for k, v := range m {
		kvSlice = append(kvSlice, kv{k, v})
	}
	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].Value > kvSlice[j].Value
	})

	res := make([]rune, 0, len(s))
	for _, kvs := range kvSlice {
		for i := 0; i < kvs.Value; i++ {
			res = append(res, kvs.Key)
		}
	}
	return string(res)
}

type kv struct {
	Key   rune
	Value int
}
