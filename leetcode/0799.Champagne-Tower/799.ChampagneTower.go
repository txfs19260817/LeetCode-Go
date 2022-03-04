package _799_Champagne_Tower

func champagneTower(poured int, query_row int, query_glass int) float64 {
	t := make([][]float64, 102)
	for i := range t {
		t[i] = make([]float64, 102)
	}
	t[0][0] = float64(poured)
	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			q := (t[r][c] - 1.0) / 2
			if q > 0 {
				t[r+1][c] += q
				t[r+1][c+1] += q
			}
		}
	}
	if t[query_row][query_glass] > 1 {
		return 1
	}
	return t[query_row][query_glass]
}
