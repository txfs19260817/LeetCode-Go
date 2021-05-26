package _482_Minimum_Number_of_Days_to_Make_m_Bouquets

func minDays(bloomDay []int, m int, k int) int {
	start, end := minmaxInSlice(bloomDay)
	for start+1 < end {
		mid := start + (end-start)/2
		if check(bloomDay, m, k, mid) {
			end = mid
		} else {
			start = mid
		}
	}
	if check(bloomDay, m, k, start) {
		return start
	}
	if check(bloomDay, m, k, end) {
		return end
	}
	return -1
}

func minmaxInSlice(a []int) (mini, maxi int) {
	mini = a[0]
	for _, n := range a {
		if n > maxi {
			maxi = n
		}
		if n < mini {
			mini = n
		}
	}
	return
}

func check(bloomDay []int, m, k, days int) bool {
	copyK := k
	for _, d := range bloomDay {
		if d <= days {
			copyK--
			if copyK == 0 {
				m--
				if m <= 0 {
					return true
				}
				copyK = k
			}
		} else {
			copyK = k
		}
	}
	return false
}
