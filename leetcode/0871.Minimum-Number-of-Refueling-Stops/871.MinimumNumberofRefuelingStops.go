package leetcode

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	dp := make([]int, len(stations)+1)
	dp[0] = startFuel
	for i := 0; i < len(stations); i++ {
		for t := i; t >= 0; t-- {
			if dp[t] >= stations[i][0] {
				dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}
	for i, n := range dp {
		if n >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minRefuelStops2(target int, startFuel int, stations [][]int) int {
	var ans, prev int
	var h intHeap
	stations = append(stations, []int{target, 0})
	for _, station := range stations {
		location, capacity := station[0], station[1]
		startFuel -= location - prev // `startFuel` represents remaining fuel in tank
		for len(h) > 0 && startFuel < 0 {
			startFuel += heap.Pop(&h).(int)
			ans++
		}
		if startFuel < 0 {
			return -1
		}
		if location == target {
			break
		}
		heap.Push(&h, capacity)
		prev = location
	}
	return ans
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
