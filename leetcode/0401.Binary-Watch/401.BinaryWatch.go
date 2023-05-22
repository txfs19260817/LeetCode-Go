package leetcode

import "strconv"

type hm struct {
	hour, min int
}

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	base := []hm{
		{1, 0}, {2, 0}, {4, 0}, {8, 0},
		{0, 1}, {0, 2}, {0, 4}, {0, 8}, {0, 16}, {0, 32},
	}
	if num == 1 {
		return format(base)
	}
	var ans []hm
	dfs(&ans, base, []hm{}, 0, num)
	return format(ans)
}

func dfs(ans *[]hm, base, path []hm, index, num int) {
	if num == 0 {
		if s := sum(path); s.min <= 59 && s.hour <= 11 {
			*ans = append(*ans, s)
		}
		return
	}
	for i := index; i < len(base); i++ {
		path = append(path, base[i])
		dfs(ans, base, path, i+1, num-1)
		path = path[:len(path)-1]
	}
}

func format(times []hm) []string {
	ans := make([]string, 0, len(times))
	for _, time := range times {
		if time.hour > 11 || time.min > 59 {
			continue
		}
		s := strconv.Itoa(time.hour) + ":"
		if time.min < 10 {
			s += "0"
		}
		s += strconv.Itoa(time.min)
		ans = append(ans, s)
	}
	return ans
}

func sum(times []hm) hm {
	var ans hm
	for _, time := range times {
		ans.hour += time.hour
		ans.min += time.min
	}
	return ans
}
