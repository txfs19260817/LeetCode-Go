package leetcode

import (
	"cmp"
	"slices"
	"sort"
)

type fenwick []int // 1-based

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}
func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(xCoord []int, yCoord []int) int64 {
	type pair struct{ x, y int }
	points := make([]pair, len(xCoord))
	for i := range xCoord {
		points[i] = pair{xCoord[i], yCoord[i]}
	}
	slices.SortFunc(points, func(a, b pair) int { return cmp.Or(a.x-b.x, a.y-b.y) })

	// y 离散化（把大坐标压到 0..m-1） - tree 只能用连续下标，所以把所有出现过的 y 排序去重，得到“字典序 y 列表”。
	slices.Sort(yCoord)
	yCoord = slices.Compact(yCoord)

	ans := int64(-1)
	tree := make(fenwick, len(yCoord)+1)               // 1-based
	tree.add(sort.SearchInts(yCoord, points[0].y) + 1) // From the first point,add its y into tree

	type tuple struct{ x, y, cnt int }
	pre := make([]tuple, len(yCoord)+1) // Tracks the previous vertical segment ending at Y. 1-based

	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1].x, points[i-1].y
		x2, y2 := points[i].x, points[i].y
		y2Idx := sort.SearchInts(yCoord, y2) + 1
		tree.add(y2Idx)

		if x1 != x2 { // Points are not in the same column
			continue
		}

		y1Idx := sort.SearchInts(yCoord, y1) + 1
		cur := tree.query(y1Idx, y2Idx) // 落在 [y1,y2] 的点共有多少
		p := pre[y2Idx]                 // p 是之前记录的“左竖边”的快照，用来和当前竖边对比增量

		// 从左竖边到右竖边之间，如果矩形内部/边界没有别的点，那么新增的只会是右边两个角点 ⇒ 区间计数只增加 2。
		// p.y==y1 确保上下端一致（同一个 y1,y2），于是面积就是宽 (x2-p.x) × 高 (y2-y1)。
		if p.cnt > 0 && p.cnt+2 == cur && p.y == y1 {
			ans = max(ans, int64(x2-p.x)*int64(y2-y1))
		}
		pre[y2Idx] = tuple{x1, y1, cur}
	}

	return ans
}
