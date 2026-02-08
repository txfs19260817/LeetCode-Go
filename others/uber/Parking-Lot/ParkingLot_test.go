package uber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParkingLotParkRemove(t *testing.T) {
	lot := NewParkingLot(1, 4, 1)

	assert.True(t, lot.Park("m1", "motorcycle"))
	assert.True(t, lot.Park("c1", "car"))
	assert.True(t, lot.Park("v1", "van"))
	assert.False(t, lot.Park("b1", "bus"))

	assert.True(t, lot.Remove("v1"))
	assert.True(t, lot.Park("b1", "bus"))
	assert.False(t, lot.Park("b1", "bus"))
	assert.False(t, lot.Remove("nope"))
}

func TestParkingLotCompactFallback(t *testing.T) {
	lot := NewParkingLot(0, 3, 0)

	assert.True(t, lot.Park("v2", "van"))
	assert.Equal(t, 0, lot.Available(SpotCompact))
	assert.True(t, lot.Remove("v2"))
	assert.Equal(t, 3, lot.Available(SpotCompact))
}
