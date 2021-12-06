package sn

func largestSubGrid(grid [][]int, maxSum int) int {
	n := len(grid)
	prefix := make([][]int, n)
	for i := range prefix {
		prefix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				prefix[0][0] = grid[0][0]
			} else if i == 0 {
				prefix[0][j] = prefix[0][j-1] + grid[0][j]
			} else if j == 0 {
				prefix[i][0] = prefix[i-1][0] + grid[i][0]
			} else {
				prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + grid[i][j]
			}
		}
	}
	l, r := 0, n
	for l < r {
		mid, largestSum := l+(r-l+1)/2, 0
		for i := mid - 1; i < n; i++ {
			for j := mid - 1; j < n; j++ {
				sum := prefix[i][j]
				if i >= mid {
					sum -= prefix[i-mid][j]
				}
				if j >= mid {
					sum -= prefix[i][j-mid]
				}
				if i >= mid && j >= mid {
					sum += prefix[i-mid][j-mid]
				}
				if largestSum < sum {
					largestSum = sum
				}
			}
		}
		if largestSum <= maxSum {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}
