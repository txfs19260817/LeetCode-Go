package snowflake

import (
	"runtime"
	"strings"
	"sync"
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

// GrepParallel is a concurrent version of Grep that uses goroutines to speed up the search.
func GrepParallel(lines []string, target string, linesAround int) []string {
	n := len(lines)
	if n == 0 {
		return nil
	}

	// 1. Parallelize the search phase (CPU bound)
	numWorkers := runtime.GOMAXPROCS(0)
	chunkSize := (n + numWorkers - 1) / numWorkers

	// Store match indices. Using a slice of slices to avoid locking on every append
	// or a channel.
	matchIndices := make([][]int, numWorkers)
	var wg sync.WaitGroup

	for i := range numWorkers {
		start := i * chunkSize
		if start >= n {
			break
		}
		end := min(start + chunkSize, n)

		wg.Add(1)
		go func(workerID, s, e int) {
			defer wg.Done()
			var localMatches []int
			for j := s; j < e; j++ {
				if strings.Contains(lines[j], target) {
					localMatches = append(localMatches, j)
				}
			}
			matchIndices[workerID] = localMatches
		}(i, start, end)
	}
	wg.Wait()

	// 2. Aggregate results and mark lines (Sequential, usually fast)
	shouldPrint := make([]bool, n)
	for _, matches := range matchIndices {
		for _, idx := range matches {
			start := max(idx-linesAround, 0)
			end := min(idx+linesAround, n-1)
			for j := start; j <= end; j++ {
				shouldPrint[j] = true
			}
		}
	}

	// 3. Collect final output
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

// GrepStreamingParallel processes a stream of lines concurrently.
// It preserves the input order for the output context logic.
func GrepStreamingParallel(input <-chan string, target string, linesAround int) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		// Internal structures for coordination
		type job struct {
			id   int
			line string
		}
		type result struct {
			id      int
			line    string
			matches bool
		}

		jobs := make(chan job, 100)
		results := make(chan result, 100)
		numWorkers := runtime.GOMAXPROCS(0)
		var wg sync.WaitGroup

		// 1. Start Workers
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := range jobs {
					matches := strings.Contains(j.line, target)
					results <- result{
						id:      j.id,
						line:    j.line,
						matches: matches,
					}
				}
			}()
		}

		// 2. Start Dispatcher (Producer)
		go func() {
			id := 0
			for line := range input {
				jobs <- job{id: id, line: line}
				id++
			}
			close(jobs)
			wg.Wait()
			close(results)
		}()

		// 3. Start Collector (Consumer & Reordering & Context Logic)
		// We reuse the logic from StreamingGrep effectively, but we must handle out-of-order results.
		sg := NewStreamingGrep(linesAround, target)
		buffer := make(map[int]result)
		nextID := 0

		for res := range results {
			// Buffer the result
			buffer[res.id] = res

			// Process as many sequential results as possible
			for {
				r, ok := buffer[nextID]
				if !ok {
					break
				}
				delete(buffer, nextID) // Clean up map

				// Custom ProcessLine logic to use pre-calculated 'matches'
				// We can't directly use sg.ProcessLine because it calls Contains again.
				// We replicate the logic here or modify sg.ProcessLine.
				// For clean code, let's just replicate the state logic here since we have `matches` boolean.

				var out []string
				if r.matches {
					out = append(out, sg.buffer...)
					sg.buffer = sg.buffer[:0]
					out = append(out, r.line)
					sg.remainingAfter = sg.linesAround
				} else if sg.remainingAfter > 0 {
					out = append(out, r.line)
					sg.remainingAfter--
				} else {
					sg.buffer = append(sg.buffer, r.line)
					if len(sg.buffer) > sg.linesAround {
						sg.buffer = sg.buffer[1:]
					}
				}

				// Send to output
				for _, l := range out {
					output <- l
				}

				nextID++
			}
		}
	}()

	return output
}
