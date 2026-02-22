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

// StreamEncoder encodes a stream of non-negative integers using RLE and BP runs.
type StreamEncoder struct {
	curVal  int
	curCnt  int
	bpBuf   []int
	started bool
}

// NewStreamEncoder creates a new StreamEncoder.
func NewStreamEncoder() *StreamEncoder {
	return &StreamEncoder{}
}

// Write feeds one integer into the encoder.
// Returns any encoded runs that are now complete (may be empty).
func (e *StreamEncoder) Write(value int) []string {
	if !e.started {
		e.started = true
		e.curVal = value
		e.curCnt = 1
		return nil
	}

	if value == e.curVal {
		e.curCnt++
		return nil
	}

	result := e.finalizeRun(false)
	e.curVal = value
	e.curCnt = 1
	return result
}

// Flush signals end of stream. Returns any remaining encoded runs.
func (e *StreamEncoder) Flush() []string {
	if !e.started {
		return nil
	}
	result := e.finalizeRun(true)
	e.started = false
	e.curCnt = 0
	return result
}

// finalizeRun processes the current run (curVal, curCnt) and flushes the BP
// buffer as needed. isLast indicates this is the final run in the stream.
func (e *StreamEncoder) finalizeRun(isLast bool) []string {
	var result []string

	if e.curCnt >= 8 {
		result = append(result, e.flushBP()...)
		result = append(result, fmt.Sprintf("RLE[%d,%d]", e.curVal, e.curCnt))
	} else if isLast && len(e.bpBuf) == 0 {
		result = append(result, fmt.Sprintf("RLE[%d,%d]", e.curVal, e.curCnt))
	} else {
		for i := 0; i < e.curCnt; i++ {
			e.bpBuf = append(e.bpBuf, e.curVal)
			if len(e.bpBuf) == 8 {
				result = append(result, e.flushBP()...)
			}
		}
		if isLast {
			result = append(result, e.flushBP()...)
		}
	}
	return result
}

// flushBP emits a BP run from the buffer and clears it.
func (e *StreamEncoder) flushBP() []string {
	if len(e.bpBuf) == 0 {
		return nil
	}
	parts := make([]string, len(e.bpBuf))
	for i, v := range e.bpBuf {
		parts[i] = strconv.Itoa(v)
	}
	e.bpBuf = e.bpBuf[:0]
	return []string{fmt.Sprintf("BP[%s]", strings.Join(parts, ","))}
}

// StreamDecoder decodes encoded run strings one at a time.
type StreamDecoder struct{}

// NewStreamDecoder creates a new StreamDecoder.
func NewStreamDecoder() *StreamDecoder {
	return &StreamDecoder{}
}

// Write feeds one encoded run string and returns the decoded integers.
func (d *StreamDecoder) Write(run string) []int {
	if strings.HasPrefix(run, "RLE[") {
		inner := run[4 : len(run)-1]
		parts := strings.Split(inner, ",")
		value, _ := strconv.Atoi(parts[0])
		count, _ := strconv.Atoi(parts[1])
		result := make([]int, count)
		for i := range result {
			result[i] = value
		}
		return result
	}
	if strings.HasPrefix(run, "BP[") {
		inner := run[3 : len(run)-1]
		parts := strings.Split(inner, ",")
		result := make([]int, len(parts))
		for i, p := range parts {
			result[i], _ = strconv.Atoi(p)
		}
		return result
	}
	return nil
}
