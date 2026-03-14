package leetcode

import "container/heap"

type pair struct {
	nextNode    int
	probability float64
}
type maxHeapq []pair

func (h maxHeapq) Len() int {
	return len(h)
}

func (h maxHeapq) Less(i, j int) bool {
	return h[i].probability > h[j].probability
}

func (h maxHeapq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeapq) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *maxHeapq) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	g := map[int][]pair{}
	for i, e := range edges {
		g[e[0]] = append(g[e[0]], pair{nextNode: e[1], probability: succProb[i]})
		g[e[1]] = append(g[e[1]], pair{nextNode: e[0], probability: succProb[i]})
	}

	hq := &maxHeapq{pair{nextNode: startNode, probability: 1.}}
	probs := make([]float64, n)
	probs[startNode] = 1.

	for hq.Len() > 0 {
		p := heap.Pop(hq).(pair)
		prb, node := p.probability, p.nextNode
		// skips stale heap entries
		if prb < probs[node] {
			continue
		}
		// visit neighbors
		for _, p2 := range g[node] {
			nextPrb, nextNode := p2.probability, p2.nextNode
			if probs[nextNode] < nextPrb*probs[node] {
				probs[nextNode] = nextPrb * probs[node]
				heap.Push(hq, pair{nextNode, probs[nextNode]})
			}
		}
	}

	return probs[endNode]
}
