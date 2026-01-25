package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeetingScheduler(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		scheduler := Constructor([]string{"roomB", "roomA", "roomC"})
		assert.Equal(t, "roomA", scheduler.Schedule(1, 5))
		assert.Equal(t, "roomB", scheduler.Schedule(1, 5))
		assert.Equal(t, "roomC", scheduler.Schedule(2, 6))
		assert.Equal(t, "", scheduler.Schedule(2, 3))
		assert.Equal(t, "roomA", scheduler.Schedule(5, 10))
		assert.Equal(t, "roomB", scheduler.Schedule(8, 10))
	})

	t.Run("Example 2", func(t *testing.T) {
		scheduler := Constructor([]string{"roomA"})
		assert.Equal(t, "roomA", scheduler.Schedule(1, 5))
		assert.Equal(t, "", scheduler.Schedule(1, 5))
		assert.Equal(t, "roomA", scheduler.Schedule(5, 10))
		assert.Equal(t, "", scheduler.Schedule(2, 6))
	})

	t.Run("Example 3", func(t *testing.T) {
		scheduler := Constructor([]string{"roomA", "roomB"})
		assert.Equal(t, "roomA", scheduler.Schedule(1, 3))
		assert.Equal(t, "roomB", scheduler.Schedule(2, 4))
		assert.Equal(t, "roomA", scheduler.Schedule(3, 5))
		assert.Equal(t, "roomB", scheduler.Schedule(1, 2))
		assert.Equal(t, "roomB", scheduler.Schedule(4, 6))
		assert.Equal(t, "roomA", scheduler.Schedule(8, 10))
	})
}
