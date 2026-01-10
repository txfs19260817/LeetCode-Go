package snowflake

import (
	"testing"
)

func TestLRUCacheOperations(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     [][]int
		expected []interface{}
	}{
		{
			name:     "Test Case 1",
			commands: []string{"LRUCache", "put", "put", "put", "get", "put", "get", "resize", "put", "put", "resize", "get", "get", "put", "get"},
			args:     [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {2}, {4, 4}, {1}, {5}, {5, 5}, {6, 6}, {2}, {5}, {4}, {6, 66}, {6}},
			expected: []interface{}{nil, nil, nil, nil, 2, nil, -1, nil, nil, nil, nil, 5, -1, nil, 66},
		},
		{
			name:     "Test Case 2",
			commands: []string{"LRUCache", "put", "put", "resize", "put", "get", "get"},
			args:     [][]int{{2}, {1, 10}, {2, 20}, {1}, {3, 30}, {1}, {3}},
			expected: []interface{}{nil, nil, nil, nil, nil, -1, 30},
		},
		{
			name:     "Test Case 3",
			commands: []string{"LRUCache", "put", "put", "resize", "get", "put", "get", "get"},
			args:     [][]int{{1}, {5, 5}, {6, 6}, {3}, {5}, {7, 7}, {6}, {7}},
			expected: []interface{}{nil, nil, nil, nil, -1, nil, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj LRUCache
			for i, cmd := range tt.commands {
				arg, want := tt.args[i], tt.expected[i]

				switch cmd {
				case "LRUCache":
					obj = NewLRUCache(arg[0])
				case "put":
					obj.Put(arg[0], arg[1])
				case "get":
					got := obj.Get(arg[0])
					if want != nil && got != want.(int) {
						t.Errorf("step %d: Get(%d) = %d, want %d", i, arg[0], got, want)
					}
				case "resize":
					obj.Resize(arg[0])
				}
			}
		})
	}
}
