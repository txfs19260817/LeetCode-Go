package waymo

import "math"

func KMeans(points [][]float64, k int, maxIter int) ([][]float64, []int) {
	centroids := cloneMatrix(points[:k])
	assignments := make([]int, len(points))

	for iter := 0; iter < maxIter; iter++ {
		assign(points, centroids, assignments)
		next := recompute(points, centroids, assignments, k)
		if sameCentroids(centroids, next) {
			centroids = next
			break
		}
		centroids = next
	}

	assign(points, centroids, assignments)
	return centroids, assignments
}

func assign(points [][]float64, centroids [][]float64, assignments []int) {
	for i, p := range points {
		bestIdx := 0
		bestDist := distanceSquared(p, centroids[0])
		for j := 1; j < len(centroids); j++ {
			d := distanceSquared(p, centroids[j])
			if d < bestDist {
				bestDist = d
				bestIdx = j
			}
		}
		assignments[i] = bestIdx
	}
}

func recompute(points [][]float64, centroids [][]float64, assignments []int, k int) [][]float64 {
	dim := len(points[0])
	next := make([][]float64, k)
	counts := make([]int, k)
	for i := 0; i < k; i++ {
		next[i] = make([]float64, dim)
	}

	for i, p := range points {
		c := assignments[i]
		for d := 0; d < dim; d++ {
			next[c][d] += p[d]
		}
		counts[c]++
	}

	for i := 0; i < k; i++ {
		if counts[i] == 0 {
			copy(next[i], centroids[i])
			continue
		}
		inv := 1.0 / float64(counts[i])
		for d := 0; d < dim; d++ {
			next[i][d] *= inv
		}
	}
	return next
}

func sameCentroids(a [][]float64, b [][]float64) bool {
	for i := range a {
		for d := range a[i] {
			if math.Abs(a[i][d]-b[i][d]) > 1e-9 {
				return false
			}
		}
	}
	return true
}

func distanceSquared(a []float64, b []float64) float64 {
	sum := 0.0
	for i := range a {
		delta := a[i] - b[i]
		sum += delta * delta
	}
	return sum
}

func cloneMatrix(src [][]float64) [][]float64 {
	out := make([][]float64, len(src))
	for i := range src {
		out[i] = make([]float64, len(src[i]))
		copy(out[i], src[i])
	}
	return out
}
