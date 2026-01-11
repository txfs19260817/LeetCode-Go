package snowflake

import (
	"strings"
)

// Grep searches for lines matching the target and returns them along with linesAround context.
// No line is printed more than once.
func Grep(lines []string, target string, linesAround int) []string {
	n := len(lines)
	shouldPrint := make([]bool, n)

	// Mark lines to be printed
	for i, line := range lines {
		if strings.Contains(line, target) {
			start := max(i-linesAround, 0)
			end := min(i+linesAround, n-1)
			for j := start; j <= end; j++ {
				shouldPrint[j] = true
			}
		}
	}

	// Collect results
	var result []string
	for i, print := range shouldPrint {
		if print {
			result = append(result, lines[i])
		}
	}
	return result
}

// StreamingGrep handles a stream of lines and prints matches with context.
type StreamingGrep struct {
	linesAround    int
	target         string
	buffer         []string // Stores lines for "before" context
	remainingAfter int      // Counts how many "after" lines we still need to print
}

// NewStreamingGrep creates a new streaming grep processor.
func NewStreamingGrep(linesAround int, target string) *StreamingGrep {
	return &StreamingGrep{
		linesAround: linesAround,
		target:      target,
		buffer:      make([]string, 0, linesAround),
	}
}

// ProcessLine processes a single line and returns lines that should be printed now.
func (s *StreamingGrep) ProcessLine(line string) []string {
	var output []string

	if strings.Contains(line, s.target) {
		// 1. Output everything in buffer (these are "before" lines)
		output = append(output, s.buffer...)
		s.buffer = s.buffer[:0] // Clear buffer as they are now printed

		// 2. Output current line
		output = append(output, line)

		// 3. Reset countdown for "after" lines
		s.remainingAfter = s.linesAround
	} else if s.remainingAfter > 0 {
		// It's an "after" line for a previous match
		output = append(output, line)
		s.remainingAfter--
	} else {
		// Potential "before" line for future match
		s.buffer = append(s.buffer, line)
		if len(s.buffer) > s.linesAround {
			// Remove oldest to maintain window size
			s.buffer = s.buffer[1:]
		}
	}

	return output
}
