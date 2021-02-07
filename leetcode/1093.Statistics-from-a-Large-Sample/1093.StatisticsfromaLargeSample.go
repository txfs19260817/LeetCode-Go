package _093_Statistics_from_a_Large_Sample

func sampleStats(count []int) []float64 {
	// minimum, maximum, mean, median, mode
	res, cnt, acc, mode := []float64{255, 0, 0, 0, 0}, 0, 0, 0
	for _, v := range count {
		cnt += v
	}
	midLeft, midRight := cnt/2, cnt/2
	if cnt%2 == 0 {
		midRight++
	}
	for i, v := range count {
		if v > 0 {
			if i < int(res[0]) {
				res[0] = float64(i)
			}
			res[1] = float64(i)
		}
		res[2] += float64(i*v) / float64(cnt)
		if acc < midLeft && acc+v >= midLeft {
			res[3] += float64(i) / 2.0
		}
		if acc < midRight && acc+v >= midRight {
			res[3] += float64(i) / 2.0
		}
		acc += v
		if mode < v {
			mode = v
			res[4] = float64(i)
		}
	}
	return res
}
