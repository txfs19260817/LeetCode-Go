package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVStoreWithQPS_Sample(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("foo", "bar", 1)
	s.Put("baz", "qux", 2)
	assert.Equal(t, "bar", s.Get("foo", 3))
	// 3 hits / 300 = 0.01
	assert.InDelta(t, 0.01, s.GetQPS(3), 1e-9)
}

func TestKVStoreWithQPS_Expiry(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("a", "1", 1)   // expires after t=300
	s.Put("b", "2", 2)   // expires after t=301
	s.Put("c", "3", 301) // fresh

	// At t=301, window is (1, 301]. t=1 expired, t=2 and t=301 survive → 2/300.
	assert.InDelta(t, 2.0/300, s.GetQPS(301), 1e-9)
}

func TestKVStoreWithQPS_AllExpired(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("a", "1", 1)
	s.Put("b", "2", 2)

	// At t=500, both hits are well outside the 300s window.
	assert.InDelta(t, 0.0, s.GetQPS(500), 1e-9)
}

func TestKVStoreWithQPS_SameTimestamp(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("a", "1", 5)
	s.Put("b", "2", 5)
	s.Put("c", "3", 5)
	// 3 hits / 300 = 0.01
	assert.InDelta(t, 0.01, s.GetQPS(5), 1e-9)
}

func TestKVStoreWithQPS_GetCountsAsHit(t *testing.T) {
	s := NewKVStoreWithQPS()
	// Get on a missing key still records a hit.
	assert.Equal(t, "", s.Get("x", 1))
	assert.InDelta(t, 1.0/300, s.GetQPS(1), 1e-9)
}

func TestKVStoreWithQPS_Overwrite(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("k", "v1", 1)
	s.Put("k", "v2", 2)
	assert.Equal(t, "v2", s.Get("k", 3))
	// 2 puts + 1 get = 3 hits → 3/300 = 0.01
	assert.InDelta(t, 0.01, s.GetQPS(3), 1e-9)
}

func TestKVStoreWithQPS_BoundaryExact300(t *testing.T) {
	s := NewKVStoreWithQPS()
	s.Put("a", "1", 1)
	// At t=300, window is (0, 300]. t=1 is inside → 1/300.
	assert.InDelta(t, 1.0/300, s.GetQPS(300), 1e-9)
	// At t=301, window is (1, 301]. t=1 is now outside → 0.
	assert.InDelta(t, 0.0, s.GetQPS(301), 1e-9)
}

func TestKVStoreWithQPS_HighThroughput(t *testing.T) {
	s := NewKVStoreWithQPS()
	// 600 puts in a single second → 600/300 = 2.0 QPS
	for i := 0; i < 600; i++ {
		s.Put("k", "v", 10)
	}
	assert.InDelta(t, 2.0, s.GetQPS(10), 1e-9)
}
