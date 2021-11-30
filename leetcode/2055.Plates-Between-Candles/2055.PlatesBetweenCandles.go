package _055_Plates_Between_Candles

import "sort"

func platesBetweenCandles(s string, queries [][]int) []int {
	ans, plates, sum, candles := make([]int, len(queries)), make([]int, len(s)), 0, make([]int, 0, 2)
	for i, c := range s {
		if c == '|' {
			candles = append(candles, i)
			if sum == 0 { // start to accumulate
				sum = 1
			}
			if i > 0 {
				plates[i] = plates[i-1]
			}
		} else if sum > 0 { // c == '*'
			plates[i] = sum
			sum++
		}
	}
	for i, query := range queries {
		l, r := query[0], query[1]
		if s[l] != '|' { // left pt finds the next candle
			idx := sort.SearchInts(candles, l)
			if idx == len(candles) { // no candle on its right
				ans[i] = 0
				continue
			}
			l = candles[idx]
		}
		if s[r] != '|' { // right pt finds the previous candle
			idx := sort.SearchInts(candles, r)
			if idx-1 < 0 { // no candle on its left
				ans[i] = 0
				continue
			}
			r = candles[idx-1]
		}
		if l >= r {
			ans[i] = 0
		} else {
			ans[i] = plates[r] - plates[l]
		}
	}
	return ans
}

func platesBetweenCandles2(s string, queries [][]int) []int {
	ans, plates := make([]int, len(queries)), make([]int, len(s)+1) // plates[i]: the num of plates in range s[:i]
	left, p := make([]int, len(s)), -1                              // left[i]: i's left plate index
	for i, b := range s {
		plates[i+1] = plates[i]
		if b == '|' {
			p = i
		} else {
			plates[i+1]++
		}
		left[i] = p
	}
	right, p := make([]int, len(s)), len(s) // right[i]: i's right plate index
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			p = i
		}
		right[i] = p
	}

	for i, q := range queries {
		if l, r := right[q[0]], left[q[1]]; l < r {
			ans[i] = plates[r] - plates[l]
		}
	}
	return ans
}
