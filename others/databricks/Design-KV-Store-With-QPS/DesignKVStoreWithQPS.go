package databricks

const qpsWindow = 300 // 5-minute window in seconds

// KVStoreWithQPS is a key-value store that tracks the number of
// operations (Put + Get) within a sliding 5-minute window.
//
// Space-optimised approach (LC 362 variant): a circular buffer of size
// 300 where each slot stores the timestamp it belongs to and the
// operation count for that second.  This gives O(W) space regardless
// of throughput.
type KVStoreWithQPS struct {
	data   map[string]string
	times  [qpsWindow]int
	counts [qpsWindow]int
}

// NewKVStoreWithQPS creates a new empty KV store.
func NewKVStoreWithQPS() *KVStoreWithQPS {
	return &KVStoreWithQPS{
		data: make(map[string]string),
	}
}

// hit records one operation at the given timestamp.
func (s *KVStoreWithQPS) hit(timestamp int) {
	idx := timestamp % qpsWindow
	if s.times[idx] == timestamp {
		s.counts[idx]++
	} else {
		s.times[idx] = timestamp
		s.counts[idx] = 1
	}
}

// Put stores a key-value pair and records a hit.
func (s *KVStoreWithQPS) Put(key, value string, timestamp int) {
	s.data[key] = value
	s.hit(timestamp)
}

// Get retrieves the value for a key (empty string if absent) and records a hit.
func (s *KVStoreWithQPS) Get(key string, timestamp int) string {
	s.hit(timestamp)
	return s.data[key]
}

// GetQPS returns the average queries-per-second over the last 300 seconds,
// i.e. (total hits in (timestamp-300, timestamp]) / 300.
func (s *KVStoreWithQPS) GetQPS(timestamp int) float64 {
	total := 0
	for i := 0; i < qpsWindow; i++ {
		if timestamp-s.times[i] < qpsWindow {
			total += s.counts[i]
		}
	}
	return float64(total) / float64(qpsWindow)
}
