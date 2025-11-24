package k

import "strconv"

func shopping(products [][]string, list []string) int {
	product2department := map[string]string{}
	for _, p := range products {
		product2department[p[0]] = p[1]
	}

	var sequential []string
	grouped := map[string]bool{}
	for _, product := range list {
		department := product2department[product] // assuming the product always exists in map
		if len(sequential) == 0 || department != sequential[len(sequential)-1] {
			sequential = append(sequential, department)
		}
		grouped[department] = true
	}

	return len(sequential) - len(grouped)
}

func carpool(roads [][]string, starts []string, people [][]string) [][]string {
	ans := make([][]string, 2)
	type Edge struct {
		to       string
		duration int
	}
	roadMap := map[string]Edge{}

	for _, r := range roads {
		// [Origin, Destination, Duration it takes to drive]
		d, _ := strconv.Atoi(r[2])
		roadMap[r[0]] = Edge{r[1], d}
	}

	loc2Duration := func(start string) map[string]int {
		l2d := map[string]int{}
		from := start
		l2d[from] = 0

		for {
			if next, ok := roadMap[from]; ok {
				l2d[next.to] = l2d[from] + next.duration
				from = next.to
			} else {
				break
			}
		}
		return l2d
	}

	time1 := loc2Duration(starts[0])
	time2 := loc2Duration(starts[1])

	for _, entry := range people {
		name, loc := entry[0], entry[1]
		dur1, ok1 := time1[loc]
		dur2, ok2 := time2[loc]

		if ok1 && ok2 {
			if dur1 < dur2 {
				ans[0] = append(ans[0], name)
			} else {
				ans[1] = append(ans[1], name) // if dur1==dur2, either car is okay
			}
		} else if ok1 {
			ans[0] = append(ans[0], name)
		} else if ok2 {
			ans[1] = append(ans[1], name)
		}
	}
	return ans
}
