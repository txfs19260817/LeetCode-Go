package uber

import (
	"fmt"
	"sort"
	"strconv"
)

func CalculateTaskTimes(logs [][]string) []string {
	type logEntry struct {
		op    string
		enter bool
		ts    int
	}
	logEntries := make([]logEntry, len(logs))
	for i, log := range logs {
		ts, _ := strconv.Atoi(log[2])
		logEntries[i] = logEntry{log[0], log[1] == "enter", ts}
	}
	sort.Slice(logEntries, func(i, j int) bool { return logEntries[i].ts < logEntries[j].ts })

	var stack []logEntry
	op2time := map[string]int{}
	for _, e := range logEntries {
		if e.enter {
			if len(stack) > 0 {
				st := stack[len(stack)-1]
				op2time[st.op] += e.ts - st.ts
			}
			stack = append(stack, e)
		} else {
			st := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2time[st.op] += e.ts - st.ts
			if len(stack) > 0 {
				stack[len(stack)-1].ts = e.ts
			}
		}
	}

	ans := make([]string, 0, len(op2time))
	for op, t := range op2time {
		ans = append(ans, fmt.Sprintf("%s: %d", op, t))
	}
	return ans
}
