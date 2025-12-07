package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSystem(t *testing.T) {
	tests := []struct {
		name     string
		methods  []string
		args     [][]string
		expected []interface{}
	}{
		{
			name:     "Example 1",
			methods:  []string{"FileSystem", "ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"},
			args:     [][]string{{}, {"/"}, {"/a/b/c"}, {"/a/b/c/d", "hello"}, {"/"}, {"/a/b/c/d"}},
			expected: []interface{}{nil, []string{}, nil, nil, []string{"a"}, "hello"},
		},
		{
			name:     "Example 2",
			methods:  []string{"FileSystem", "ls", "mkdir", "ls", "mkdir", "ls"},
			args:     [][]string{{}, {"/"}, {"/a/b/c"}, {"/a/b"}, {"/a/b/a"}, {"/a/b"}},
			expected: []interface{}{nil, []string{}, nil, []string{"c"}, nil, []string{"a", "c"}},
		},
		{
			name:     "Example 3",
			methods:  []string{"FileSystem", "mkdir", "ls", "ls", "mkdir", "ls", "ls", "addContentToFile", "ls", "ls", "ls"},
			args:     [][]string{{}, {"/goowmfn"}, {"/goowmfn"}, {"/"}, {"/z"}, {"/"}, {"/"}, {"/goowmfn/c", "shetopcy"}, {"/z"}, {"/goowmfn/c"}, {"/goowmfn"}},
			expected: []interface{}{nil, nil, []string{}, []string{"goowmfn"}, nil, []string{"goowmfn", "z"}, []string{"goowmfn", "z"}, nil, []string{}, []string{"c"}, []string{"c"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj FileSystem
			for i, method := range tt.methods {
				switch method {
				case "FileSystem":
					obj = Constructor()
				case "ls":
					got := obj.Ls(tt.args[i][0])
					want := tt.expected[i].([]string)
					assert.Equal(t, got, want)
				case "mkdir":
					obj.Mkdir(tt.args[i][0])
				case "addContentToFile":
					obj.AddContentToFile(tt.args[i][0], tt.args[i][1])
				case "readContentFromFile":
					got := obj.ReadContentFromFile(tt.args[i][0])
					want := tt.expected[i].(string)
					assert.Equal(t, got, want)
				}
			}
		})
	}
}
