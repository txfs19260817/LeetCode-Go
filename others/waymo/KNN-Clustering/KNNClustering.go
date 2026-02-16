package waymo

import "sort"

type Point struct {
	X float64
	Y float64
}

type LabeledPoint struct {
	Point Point
	Label int
}

type neighbor struct {
	label int
	dist2 float64
}

type labelStat struct {
	label     int
	votes     int
	totalDist float64
}

func AssignClusters(train []LabeledPoint, queries []Point, k int) []int {
	result := make([]int, len(queries))
	if len(train) == 0 || k <= 0 {
		for i := range result {
			result[i] = -1
		}
		return result
	}
	if k > len(train) {
		k = len(train)
	}

	for i, query := range queries {
		neighbors := make([]neighbor, len(train))
		for j, sample := range train {
			neighbors[j] = neighbor{
				label: sample.Label,
				dist2: distanceSquared(query, sample.Point),
			}
		}
		sort.Slice(neighbors, func(a, b int) bool {
			if neighbors[a].dist2 != neighbors[b].dist2 {
				return neighbors[a].dist2 < neighbors[b].dist2
			}
			return neighbors[a].label < neighbors[b].label
		})

		stats := make([]labelStat, 0, k)
		for j := 0; j < k; j++ {
			stats = addVote(stats, neighbors[j].label, neighbors[j].dist2)
		}
		result[i] = pickBestLabel(stats)
	}
	return result
}

func addVote(stats []labelStat, label int, dist2 float64) []labelStat {
	for i := range stats {
		if stats[i].label == label {
			stats[i].votes++
			stats[i].totalDist += dist2
			return stats
		}
	}
	return append(stats, labelStat{
		label:     label,
		votes:     1,
		totalDist: dist2,
	})
}

func pickBestLabel(stats []labelStat) int {
	best := stats[0]
	for i := 1; i < len(stats); i++ {
		cur := stats[i]
		if cur.votes > best.votes ||
			(cur.votes == best.votes && cur.totalDist < best.totalDist) ||
			(cur.votes == best.votes && cur.totalDist == best.totalDist && cur.label < best.label) {
			best = cur
		}
	}
	return best.label
}

func distanceSquared(a Point, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return dx*dx + dy*dy
}
