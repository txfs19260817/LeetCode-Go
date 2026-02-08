package databricks

import (
	"fmt"
	"strconv"
	"strings"
)

// Encode encodes a slice of non-negative integers using RLE and BP runs.
func Encode(values []int) []string {
	if len(values) == 0 {
		return nil
	}

	// Identify maximal runs of consecutive identical values.
	type run struct {
		value int
		count int
	}
	var runs []run
	i := 0
	for i < len(values) {
		val := values[i]
		cnt := 1
		for i+cnt < len(values) && values[i+cnt] == val {
			cnt++
		}
		runs = append(runs, run{value: val, count: cnt})
		i += cnt
	}

	var result []string
	var bpBuf []int

	flushBP := func() {
		if len(bpBuf) == 0 {
			return
		}
		parts := make([]string, len(bpBuf))
		for j, v := range bpBuf {
			parts[j] = strconv.Itoa(v)
		}
		result = append(result, fmt.Sprintf("BP[%s]", strings.Join(parts, ",")))
		bpBuf = bpBuf[:0]
	}

	for idx, r := range runs {
		isLast := idx == len(runs)-1

		if r.count >= 8 {
			// RLE-eligible: flush BP buffer first, then emit RLE.
			flushBP()
			result = append(result, fmt.Sprintf("RLE[%d,%d]", r.value, r.count))
		} else if isLast && len(bpBuf) == 0 {
			// Last run exception: emit as RLE even though count < 8.
			result = append(result, fmt.Sprintf("RLE[%d,%d]", r.value, r.count))
		} else {
			// Add values to BP buffer.
			for j := 0; j < r.count; j++ {
				bpBuf = append(bpBuf, r.value)
				if len(bpBuf) == 8 {
					flushBP()
				}
			}
		}
	}
	flushBP()

	return result
}

// Decode decodes a slice of RLE/BP run strings back into integers.
func Decode(runs []string) []int {
	var result []int
	for _, s := range runs {
		if strings.HasPrefix(s, "RLE[") {
			// "RLE[value,count]"
			inner := s[4 : len(s)-1]
			parts := strings.Split(inner, ",")
			value, _ := strconv.Atoi(parts[0])
			count, _ := strconv.Atoi(parts[1])
			for i := 0; i < count; i++ {
				result = append(result, value)
			}
		} else if strings.HasPrefix(s, "BP[") {
			// "BP[v1,v2,...,vk]"
			inner := s[3 : len(s)-1]
			parts := strings.Split(inner, ",")
			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				result = append(result, v)
			}
		}
	}
	return result
}
