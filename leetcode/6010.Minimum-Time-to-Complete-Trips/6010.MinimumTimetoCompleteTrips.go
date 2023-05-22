package leetcode

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	if len(time) == 1 {
		return int64(time[0] * totalTrips)
	}
	sort.Ints(time)
	var cnt int
	l, r := time[0], time[len(time)-1]*totalTrips
	for l+1 < r {
		t := (l + r) / 2
		cnt = 0
		for i := 0; i < len(time) && time[i] <= t; i++ {
			cnt += t / time[i]
		}
		if cnt >= totalTrips {
			r = t
		} else {
			l = t
		}
	}
	cnt = 0
	for i := 0; i < len(time) && time[i] <= l; i++ {
		cnt += l / time[i]
	}
	if cnt >= totalTrips {
		return int64(l)
	}
	return int64(r)
}
