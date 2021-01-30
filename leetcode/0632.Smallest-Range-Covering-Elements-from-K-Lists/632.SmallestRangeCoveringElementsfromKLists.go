package _632_Smallest_Range_Covering_Elements_from_K_Lists

import (
	"math"
	"sort"
)

type numIdx struct {
	val, idx int
}

func smallestRange(nums [][]int) []int {
	numIdxs := make([]numIdx, 0, len(nums)*len(nums[0]))
	for i, arr := range nums {
		for _, v := range arr {
			numIdxs = append(numIdxs, numIdx{v, i})
		}
	}
	sort.Slice(numIdxs, func(i, j int) bool {
		return numIdxs[i].val < numIdxs[j].val
	})
	res, minLen, count, freq := make([]int, 2), math.MaxInt64, 0, map[int]int{} // idx:freq
	for l, r := 0, -1; l < len(numIdxs); {
		if r+1 < len(numIdxs) && count < len(nums) {
			r++
			if freq[numIdxs[r].idx] == 0 {
				count++
			}
			freq[numIdxs[r].idx]++
		} else {
			freq[numIdxs[l].idx]--
			if freq[numIdxs[l].idx] == 0 {
				count--
			}
			l++
		}
		if count == len(nums) && numIdxs[r].val-numIdxs[l].val < minLen {
			minLen = numIdxs[r].val - numIdxs[l].val
			res[0], res[1] = numIdxs[l].val, numIdxs[r].val
		}
	}
	return res
}
