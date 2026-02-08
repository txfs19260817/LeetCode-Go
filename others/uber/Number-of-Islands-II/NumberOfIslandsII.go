package uber

// NumberOfIslandsII returns island counts after each add-land operation.
func NumberOfIslandsII(m, n int, positions [][]int) []int {
	counts, _ := numberOfIslandsIIMaxSize(m, n, positions)
	return counts
}

// NumberOfIslandsIIMaxSize returns island counts and max island size after each operation.
func NumberOfIslandsIIMaxSize(m, n int, positions [][]int) ([]int, []int) {
	counts, maxSizes := numberOfIslandsIIMaxSize(m, n, positions)
	return counts, maxSizes
}

const WATER = -1

type UnionFind struct {
	parent         []int
	size           []int
	count, maxSize int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = WATER
	}
	return &UnionFind{parent: parent, size: make([]int, n)}
}

func (uf *UnionFind) IsLand(p int) bool {
	return uf.parent[p] != WATER
}

func (uf *UnionFind) AddLand(p int) {
	if uf.IsLand(p) {
		return
	}
	uf.parent[p] = p
	uf.size[p] = 1
	uf.count++
	uf.maxSize = max(uf.maxSize, 1)
}

func (uf *UnionFind) Find(p int) int {
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
}

func (uf *UnionFind) Union(p, q int) {
	p, q = uf.Find(p), uf.Find(q)
	if p == q {
		return
	}
	if uf.size[p] < uf.size[q] {
		p, q = q, p
	}
	uf.parent[q] = p
	uf.size[p] += uf.size[q]
	uf.count--
	uf.maxSize = max(uf.maxSize, uf.size[p])
}

func numberOfIslandsIIMaxSize(m, n int, positions [][]int) (counts []int, maxSizes []int) {
	uf := NewUnionFind(m * n)
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, pos := range positions {
		i, j := pos[0], pos[1]
		uf.AddLand(i*n + j)
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && uf.IsLand(ni*n+nj) {
				uf.Union(ni*n+nj, i*n+j)
			}
		}
		counts = append(counts, uf.count)
		maxSizes = append(maxSizes, uf.maxSize)
	}
	return
}
