package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_assignTasks(t *testing.T) {
	tests := []struct {
		name    string
		servers []int
		tasks   []int
		want    []int
	}{
		{
			name:    "example 1",
			servers: []int{3, 3, 2},
			tasks:   []int{1, 2, 3, 2, 1, 2},
			want:    []int{2, 2, 0, 2, 1, 2},
		},
		{
			name:    "example 2",
			servers: []int{5, 1, 4, 3, 2},
			tasks:   []int{2, 1, 2, 4, 5, 2, 1},
			want:    []int{1, 4, 1, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, assignTasks(tt.servers, tt.tasks))
		})
	}
}
