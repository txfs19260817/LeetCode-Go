package _973_K_Closest_Points_to_Origin

func kClosest(points [][]int, K int) [][]int {
	pts := make([]PT, len(points))
	for i, point := range points {
		pts[i] = PT{point, point[0]*point[0] + point[1]*point[1]}
	}
	qs(pts, 0, len(points)-1, K)
	for i, pt := range pts[:K] {
		points[i] = pt.point
	}
	return points[:K]
}

func qs(pts []PT, l, r, K int) {
	if l == r {
		return
	}
	pivot := partition(pts, l, r)
	if pivot == K {
		return
	}
	if pivot > K {
		qs(pts, l, pivot-1, K)
	} else {
		qs(pts, pivot+1, r, K)
	}
}

func partition(pts []PT, l, r int) int {
	pivotElem := pts[r]
	i := l - 1
	for j := l; j < r; j++ {
		if pts[j].dist < pivotElem.dist {
			i++
			pts[j], pts[i] = pts[i], pts[j]
		}
	}
	i++
	pts[r], pts[i] = pts[i], pts[r]
	return i
}

type PT struct {
	point []int
	dist  int
}
