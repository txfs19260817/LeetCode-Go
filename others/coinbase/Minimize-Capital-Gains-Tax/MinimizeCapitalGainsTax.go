package minimizecapitalgainstax

import (
	"container/heap"
	"strconv"
)

type priceHeap []int

func (h priceHeap) Len() int           { return len(h) }
func (h priceHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h priceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *priceHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *priceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type Solution struct{}

func (s *Solution) CalculateMinimalTax(transactions [][]string) float64 {
	amounts := make(map[int]int)
	h := &priceHeap{}
	heap.Init(h)
	var tax float64

	for _, tx := range transactions {
		if len(tx) < 4 {
			continue
		}
		txType := tx[1]
		amount, _ := strconv.Atoi(tx[2])
		price, _ := strconv.Atoi(tx[3])

		if txType == "buy" {
			if amounts[price] == 0 {
				heap.Push(h, price)
			}
			amounts[price] += amount
			continue
		}
		if txType != "sell" {
			continue
		}

		remaining := amount
		for remaining > 0 {
			for h.Len() > 0 && amounts[(*h)[0]] == 0 {
				heap.Pop(h)
			}
			if h.Len() == 0 {
				break
			}
			buyPrice := (*h)[0]
			available := amounts[buyPrice]
			take := available
			if take > remaining {
				take = remaining
			}

			profitPerUnit := price - buyPrice
			if profitPerUnit > 0 {
				tax += float64(int64(profitPerUnit)*int64(take)) / 10.0
			}

			amounts[buyPrice] -= take
			if amounts[buyPrice] == 0 {
				heap.Pop(h)
			}
			remaining -= take
		}
	}

	return tax
}
