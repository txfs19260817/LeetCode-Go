package databricks

import (
	"fmt"
	"strconv"
	"strings"
)

// StreamEncoder encodes a stream of non-negative integers using RLE and BP runs.
type StreamEncoder struct {
	curVal  int // value of the run being built
	curCnt  int // length of the run being built
	bpBuf   []int
	started bool // whether we've received at least one value
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

	// Current run ended — finalize it, then start a new one.
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
		// RLE-eligible: flush BP buffer first, then emit RLE.
		result = append(result, e.flushBP()...)
		result = append(result, fmt.Sprintf("RLE[%d,%d]", e.curVal, e.curCnt))
	} else if isLast && len(e.bpBuf) == 0 {
		// Last run exception: emit as RLE even though count < 8.
		result = append(result, fmt.Sprintf("RLE[%d,%d]", e.curVal, e.curCnt))
	} else {
		// Add values to BP buffer.
		for i := 0; i < e.curCnt; i++ {
			e.bpBuf = append(e.bpBuf, e.curVal)
			if len(e.bpBuf) == 8 {
				result = append(result, e.flushBP()...)
			}
		}
		// On the last run, flush any remaining BP values.
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
