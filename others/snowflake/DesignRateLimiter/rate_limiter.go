package snowflake

import (
	"sync"
	"time"
)

type RateLimiter struct {
	ttl   time.Duration
	limit int

	timeQueue []time.Time
	mu        sync.Mutex
}

// NewRateLimiter creates a rate limiter with TTL in ms
func NewRateLimiter(ttl, limit int) *RateLimiter {
	if ttl <= 0 {
		panic("ttl must be greater than zero")
	}
	if limit <= 0 {
		panic("limit must be greater than zero")
	}

	return &RateLimiter{ttl: time.Duration(ttl) * time.Millisecond, limit: limit, timeQueue: make([]time.Time, 0, limit)}
}

func (rl *RateLimiter) AllowRequest() bool {
	now := time.Now()
	cutOff := now.Add(-rl.ttl)
	var head int

	rl.mu.Lock()
	defer rl.mu.Unlock()

	for head < len(rl.timeQueue) && !rl.timeQueue[head].After(cutOff) {
		head++
	}

	if head > 0 {
		if head >= len(rl.timeQueue) {
			rl.timeQueue = rl.timeQueue[:0]
		} else {
			rl.timeQueue = append([]time.Time{}, rl.timeQueue[head:]...)
		}
	}
	head = 0

	if len(rl.timeQueue) < rl.limit {
		rl.timeQueue = append(rl.timeQueue, now)
		return true
	}
	return false
}
