package leetcode

import (
	"slices"
	"sort"
)

type Fenwick struct {
	n    int
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{n: n, tree: make([]int, n+1)}
}

func (f *Fenwick) Add(i, delta int) {
	for ; i <= f.n; i += i & -i {
		f.tree[i] += delta
	}
}

func (f *Fenwick) Sum(i int) (res int) {
	for ; i > 0; i -= i & -i {
		res += f.tree[i]
	}
	return
}

func kthSmallestSubarraySum(nums []int, k int) int {
	// 1) 前缀和
	n := len(nums)
	pre := make([]int, n+1)
	var total int
	minVal := nums[0]
	for i, x := range nums {
		pre[i+1] = pre[i] + x
		total += x
		minVal = min(minVal, x)
	}

	// 2) 坐标压缩：只压缩 prefix 自己就够了
	vals := slices.Clone(pre)
	slices.Sort(vals)
	vals = slices.Compact(vals)

	// 3) 统计有多少个子数组和 <= x
	coundLe := func(mid int) (ans int) {
		f := NewFenwick(len(vals))
		var inserted int

		for _, s := range pre {
			threshold := s - mid                    // 当前把 s 当成 pre[j] 想统计前面多少个 pre[i] >= s - mid
			pos := sort.SearchInts(vals, threshold) // pos = 第一个 >= t 的位置（0-based）
			lessCnt := f.Sum(pos)                   // 前面有多少个前缀和 < t
			ans += inserted - lessCnt               // inserted = 当前之前一共插入了多少个前缀和, 所以 >= t 的个数 = inserted - lessCnt
			idx := sort.SearchInts(vals, s) + 1     // 然后把当前 s 自己插进去，供后面的 j 使用
			f.Add(idx, 1)
			inserted++
		}
		return
	}

	l, r := minVal, total
	for l <= r {
		mid := l + (r-l)/2
		if coundLe(mid) < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
