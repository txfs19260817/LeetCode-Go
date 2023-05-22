package leetcode

func pushDominoes(dominoes string) string {
	dominoes, res := "L"+dominoes+"R", make([]byte, 0, len(dominoes))
	for i, j := 0, 1; j < len(dominoes); j++ {
		if dominoes[j] == '.' {
			continue
		}
		if i > 0 {
			res = append(res, dominoes[i])
		}

		interval := j - i - 1
		if dominoes[i] == dominoes[j] {
			for k := 0; k < interval; k++ {
				res = append(res, dominoes[i])
			}
		} else if dominoes[i] == 'L' && dominoes[j] == 'R' {
			for k := 0; k < interval; k++ {
				res = append(res, '.')
			}
		} else {
			for k := 0; k < interval/2; k++ {
				res = append(res, 'R')
			}
			for k := 0; k < interval%2; k++ {
				res = append(res, '.')
			}
			for k := 0; k < interval/2; k++ {
				res = append(res, 'L')
			}
		}
		i = j
	}
	return string(res)
}
