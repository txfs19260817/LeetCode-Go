package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaskTimes(t *testing.T) {
	tests := []struct {
		name string
		logs [][]string
		want []string
	}{
		{
			name: "Example 1",
			logs: [][]string{
				{"print", "enter", "10"},
				{"malloc", "enter", "12"},
				{"malloc", "exit", "14"},
				{"write", "enter", "16"},
				{"write", "exit", "18"},
				{"write", "enter", "20"},
				{"write", "exit", "22"},
				{"print", "exit", "24"},
			},
			want: []string{"malloc: 2", "print: 8", "write: 4"},
		},
		{
			name: "Example 2",
			logs: [][]string{
				{"task1", "enter", "0"},
				{"task3", "exit", "6"},
				{"task2", "exit", "8"},
				{"task2", "enter", "2"},
				{"task3", "enter", "4"},
				{"task1", "exit", "10"},
			},
			want: []string{"task1: 4", "task2: 4", "task3: 2"},
		},
		{
			name: "Example 3",
			logs: [][]string{
				{"taskA", "enter", "0"},
				{"taskA", "exit", "5"},
				{"taskA", "enter", "6"},
				{"taskA", "exit", "10"},
				{"taskB", "enter", "10"},
				{"taskB", "exit", "15"},
			},
			want: []string{"taskA: 9", "taskB: 5"},
		},
		{
			name: "Quick Task",
			logs: [][]string{
				{"quickTask", "enter", "0"},
				{"quickTask", "exit", "1"},
				{"quickTask", "enter", "2"},
				{"quickTask", "exit", "3"},
			},
			want: []string{"quickTask: 2"},
		},
		{
			name: "Complex Task",
			logs: [][]string{
				{"task1", "enter", "0"},
				{"task2", "enter", "5"},
				{"task3", "enter", "10"},
				{"task3", "exit", "12"},
				{"task2", "exit", "15"},
				{"task1", "exit", "16"},
				{"task2", "enter", "18"},
				{"task2", "exit", "25"},
			},
			want: []string{"task1: 6", "task2: 15", "task3: 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, CalculateTaskTimes(tt.logs))
		})
	}
}
