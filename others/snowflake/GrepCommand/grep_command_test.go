package snowflake

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		target      string
		linesAround int
		expected    []string
	}{
		{
			name: "Example case",
			lines: []string{
				"good morning",
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
			target:      "Alex",
			linesAround: 1,
			expected: []string{
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
		},
		{
			name:        "No match",
			lines:       []string{"a", "b", "c"},
			target:      "z",
			linesAround: 1,
			expected:    nil,
		},
		{
			name: "Overlapping matches",
			lines: []string{
				"A",
				"Target 1",
				"B",
				"Target 2",
				"C",
			},
			target:      "Target",
			linesAround: 1,
			expected: []string{
				"A",
				"Target 1",
				"B",
				"Target 2",
				"C",
			},
		},
		{
			name: "Adjacent matches",
			lines: []string{
				"A",
				"Target 1",
				"Target 2",
				"B",
			},
			target:      "Target",
			linesAround: 1,
			expected: []string{
				"A",
				"Target 1",
				"Target 2",
				"B",
			},
		},
		{
			name:        "Match at start",
			lines:       []string{"Target", "A", "B"},
			target:      "Target",
			linesAround: 1,
			expected:    []string{"Target", "A"},
		},
		{
			name:        "Match at end",
			lines:       []string{"A", "B", "Target"},
			target:      "Target",
			linesAround: 1,
			expected:    []string{"B", "Target"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Grep(tt.lines, tt.target, tt.linesAround)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Grep() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGrepParallel(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		target      string
		linesAround int
		expected    []string
	}{
		{
			name: "Example case",
			lines: []string{
				"good morning",
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
			target:      "Alex",
			linesAround: 1,
			expected: []string{
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
		},
		{
			name:        "No match",
			lines:       []string{"a", "b", "c"},
			target:      "z",
			linesAround: 1,
			expected:    nil,
		},
		{
			name: "Large chunks test",
			lines: func() []string {
				l := make([]string, 100)
				for i := range l {
					l[i] = "none"
				}
				l[10] = "Target"
				l[50] = "Target"
				l[90] = "Target"
				return l
			}(),
			target:      "Target",
			linesAround: 1,
			expected:    []string{"none", "Target", "none", "none", "Target", "none", "none", "Target", "none"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GrepParallel(tt.lines, tt.target, tt.linesAround)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GrepParallel() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestStreamingGrep(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		target      string
		linesAround int
		expected    []string
	}{
		{
			name: "Example case streaming",
			lines: []string{
				"good morning",
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
			target:      "Alex",
			linesAround: 1,
			expected: []string{
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
		},
		{
			name: "Overlapping matches streaming",
			lines: []string{
				"A",
				"Target 1",
				"B",
				"Target 2",
				"C",
			},
			target:      "Target",
			linesAround: 1,
			expected: []string{
				"A",
				"Target 1",
				"B",
				"Target 2",
				"C",
			},
		},
		{
			name: "Gap between matches",
			lines: []string{
				"A",
				"Target",
				"B",
				"C",
				"Target",
				"D",
			},
			target:      "Target",
			linesAround: 1,
			expected: []string{
				"A",
				"Target",
				"B",
				"C",
				"Target",
				"D",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := NewStreamingGrep(tt.linesAround, tt.target)
			var got []string
			for _, line := range tt.lines {
				out := sg.ProcessLine(line)
				got = append(got, out...)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("StreamingGrep = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGrepStreamingParallel(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		target      string
		linesAround int
		expected    []string
	}{
		{
			name: "Example case streaming parallel",
			lines: []string{
				"good morning",
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
			target:      "Alex",
			linesAround: 1,
			expected: []string{
				"hello there",
				"my name is Alex",
				"my friend is albert",
				"it is nice to meet you Alex",
			},
		},
		{
			name: "Gap between matches parallel",
			lines: []string{
				"A",
				"Target",
				"B",
				"C",
				"Target",
				"D",
			},
			target:      "Target",
			linesAround: 1,
			expected: []string{
				"A",
				"Target",
				"B",
				"C",
				"Target",
				"D",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan string)
			go func() {
				for _, line := range tt.lines {
					input <- line
				}
				close(input)
			}()

			outChan := GrepStreamingParallel(input, tt.target, tt.linesAround)
			var got []string
			for line := range outChan {
				got = append(got, line)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GrepStreamingParallel = %v, want %v", got, tt.expected)
			}
		})
	}
}
