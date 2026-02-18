package generaterandomnft

import (
	"container/heap"
	"math"
	"math/rand"
	"sort"
)

type TraitValue struct {
	Name   string
	Weight int
}

type Config struct {
	Name   string
	Size   string
	Traits map[string][]TraitValue
}

type Solution struct{}

type combo struct {
	values map[string]string
	weight float64
}

type pick struct {
	key    float64
	values map[string]string
}

type pickHeap []pick

func (h pickHeap) Len() int {
	return len(h)
}

func (h pickHeap) Less(i, j int) bool {
	return h[i].key < h[j].key
}

func (h pickHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pickHeap) Push(x interface{}) {
	*h = append(*h, x.(pick))
}

func (h *pickHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// GenerateNFT covers all README stages:
// 1) Base problem: supports plain trait values (equivalent to weight 1).
// 2) Follow-up 1: output is unique; panic when n exceeds unique combinations.
// 3) Follow-up 2: supports weighted trait values and weighted unique sampling.
func (s *Solution) GenerateNFT(config Config, n int) []map[string]string {
	if n <= 0 {
		return []map[string]string{}
	}

	// Follow-up 1: compute upper bound of unique combinations.
	traitNames := make([]string, 0, len(config.Traits))
	for trait := range config.Traits {
		traitNames = append(traitNames, trait)
	}
	sort.Strings(traitNames)

	totalCombos := int64(1)
	for _, trait := range traitNames {
		values := config.Traits[trait]
		if len(values) == 0 {
			panic("trait has no values")
		}
		totalCombos *= int64(len(values))
	}
	if int64(n) > totalCombos {
		panic("n exceeds number of unique combinations")
	}

	if len(traitNames) == 0 {
		return []map[string]string{{}}
	}

	if totalCombos > int64(^uint(0)>>1) {
		panic("too many combinations to materialize")
	}

	// Follow-up 2: materialize combinations and accumulate product weights.
	// Base problem is the same path with all weights = 1.
	combos := make([]combo, 0, int(totalCombos))
	current := make(map[string]string, len(traitNames))
	var build func(idx int, weight float64)
	build = func(idx int, weight float64) {
		if idx == len(traitNames) {
			clone := make(map[string]string, len(current))
			for key, value := range current {
				clone[key] = value
			}
			combos = append(combos, combo{values: clone, weight: weight})
			return
		}

		trait := traitNames[idx]
		for _, option := range config.Traits[trait] {
			current[trait] = option.Name
			optionWeight := option.Weight
			if optionWeight <= 0 {
				optionWeight = 1
			}
			build(idx+1, weight*float64(optionWeight))
		}
		delete(current, trait)
	}
	build(0, 1)

	// Follow-up 1: if all unique combos are requested, just shuffle and return.
	if int64(n) == totalCombos {
		rand.Shuffle(len(combos), func(i, j int) {
			combos[i], combos[j] = combos[j], combos[i]
		})
		out := make([]map[string]string, 0, n)
		for i := 0; i < n; i++ {
			out = append(out, combos[i].values)
		}
		return out
	}

	// Follow-up 2: weighted random sampling without replacement.
	// Uses Efraimidis-Spirakis key sampling over the full combination set.
	h := &pickHeap{}
	heap.Init(h)
	for _, item := range combos {
		weight := item.weight
		if weight <= 0 {
			weight = 1
		}
		key := math.Pow(rand.Float64(), 1.0/weight)
		if h.Len() < n {
			heap.Push(h, pick{key: key, values: item.values})
			continue
		}
		if key > (*h)[0].key {
			(*h)[0] = pick{key: key, values: item.values}
			heap.Fix(h, 0)
		}
	}

	out := make([]map[string]string, 0, n)
	for h.Len() > 0 {
		out = append(out, heap.Pop(h).(pick).values)
	}
	return out
}
