package databricks

import (
	"sync"
	"time"
)

// KVStore is a production-style key-value store with real-time QPS tracking.
//
// Differences from KVStoreWithQPS (the interview/testable version):
//   - Uses time.Now() instead of injected timestamps.
//   - Maintains a running total so QPS() is O(1) instead of O(W).
//   - Uses min(elapsed, windowSec) as denominator to handle startup.
//   - Thread-safe via sync.Mutex.
//   - Configurable window size.
type KVStore struct {
	mu        sync.Mutex
	data      map[string]string
	windowSec int64

	buckets  []int64 // per-second operation counts (circular buffer)
	total    int64   // running sum of ops within the window
	lastSec  int64   // last second that normalize advanced to
	startSec int64   // creation second, for short-window denominator
}

// NewKVStore creates a KVStore with the given QPS window in seconds.
func NewKVStore(windowSeconds int64) *KVStore {
	nowSec := time.Now().Unix()
	return &KVStore{
		data:      make(map[string]string),
		windowSec: windowSeconds,
		buckets:   make([]int64, windowSeconds),
		lastSec:   nowSec,
		startSec:  nowSec,
	}
}

// normalize advances the internal clock to nowSec, clearing stale buckets
// and adjusting the running total.  Amortised O(1) per call.
func (s *KVStore) normalize(nowSec int64) {
	if nowSec <= s.lastSec {
		return
	}
	diff := nowSec - s.lastSec
	if diff >= s.windowSec {
		// Everything expired — fast-clear.
		for i := range s.buckets {
			s.buckets[i] = 0
		}
		s.total = 0
		s.lastSec = nowSec
		return
	}
	// Only clear the newly-passed slots.
	for sec := s.lastSec + 1; sec <= nowSec; sec++ {
		idx := sec % s.windowSec
		s.total -= s.buckets[idx]
		s.buckets[idx] = 0
	}
	s.lastSec = nowSec
}

// recordOp normalizes time and records one operation at nowSec.
func (s *KVStore) recordOp(nowSec int64) {
	s.normalize(nowSec)
	idx := nowSec % s.windowSec
	s.buckets[idx]++
	s.total++
}

// Put stores a key-value pair and records a hit.
func (s *KVStore) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.recordOp(time.Now().Unix())
	s.data[key] = value
}

// Get retrieves the value for a key and records a hit.
func (s *KVStore) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.recordOp(time.Now().Unix())
	value, ok := s.data[key]
	return value, ok
}

// QPS returns the average queries-per-second over the last windowSec seconds.
// If the store has been alive for less than windowSec, uses actual elapsed
// time as the denominator so the metric is not artificially diluted.
// Note: QPS() itself is not counted as an operation.
func (s *KVStore) QPS() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	nowSec := time.Now().Unix()
	s.normalize(nowSec)

	elapsed := nowSec - s.startSec + 1
	denom := s.windowSec
	if elapsed < denom {
		denom = elapsed
	}
	if denom <= 0 {
		return 0
	}
	return float64(s.total) / float64(denom)
}
