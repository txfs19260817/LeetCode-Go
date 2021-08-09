package _632_Rank_Transform_of_a_Matrix

import "sort"

type coord struct {
	i, j int
}

type DSU struct {
	parent, size map[int]int
}

func (d *DSU) add(x int) {
	d.parent[x], d.size[x] = x, 1
}

func (d *DSU) find(x int) int {
	if x != d.parent[x] {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(x, y int) {
	x, y = d.find(x), d.find(y)
	if d.size[y] < d.size[x] {
		y, x = x, y
	}
	d.size[y] += d.size[x]
	d.parent[x] = y
}

func (d *DSU) group() [][]int {
	root2family := map[int][]int{}
	for k := range d.parent {
		root := d.find(k)
		root2family[root] = append(root2family[root], k)
	}
	res := make([][]int, 0, len(root2family))
	for _, list := range root2family {
		res = append(res, list)
	}
	return res
}

func matrixRankTransform(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	rank, el2coords, elDistinct := make([]int, n+m), map[int][]coord{}, make([]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			el2coords[matrix[i][j]] = append(el2coords[matrix[i][j]], coord{i, j})
		}
	}
	for k := range el2coords {
		elDistinct = append(elDistinct, k)
	}
	sort.Ints(elDistinct)
	for _, x := range elDistinct {
		dsu := &DSU{parent: map[int]int{}, size: map[int]int{}}
		for _, c := range el2coords[x] {
			dsu.add(c.i)
			dsu.add(c.j + n)
		}
		for _, c := range el2coords[x] {
			dsu.union(c.i, c.j+n)
		}
		for _, group := range dsu.group() {
			var maxRank int
			for _, location := range group {
				if maxRank < rank[location] {
					maxRank = rank[location]
				}
			}
			for _, location := range group {
				rank[location] = maxRank + 1
			}
		}
		for _, c := range el2coords[x] {
			matrix[c.i][c.j] = rank[c.i]
		}
	}
	return matrix
}
