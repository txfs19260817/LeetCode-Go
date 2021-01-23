package _451_Sort_Characters_By_Frequency

import "sort"

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
