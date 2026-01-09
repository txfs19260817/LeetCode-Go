package snowflake

import (
	"sync"
	"time"
)

// multiRule represents a single rate limiting rule
type multiRule struct {
	ttl   time.Duration
	limit int
	queue []time.Time
}

// MultiRuleRateLimiter manages multiple rate limiting rules
type MultiRuleRateLimiter struct {
	mu    sync.Mutex
	rules []*multiRule
}

// NewMultiRuleRateLimiter initializes an empty rate limiter with no rules.
func NewMultiRuleRateLimiter() *MultiRuleRateLimiter {
	return &MultiRuleRateLimiter{
		rules: make([]*multiRule, 0),
	}
}

// RegisterRule adds a new rate limiting rule with the specified ttl and limit.
// Returns the rule ID (index) for future reference.
func (m *MultiRuleRateLimiter) RegisterRule(ttl, limit int) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.rules = append(m.rules, &multiRule{
		ttl:   time.Duration(ttl) * time.Millisecond,
		limit: limit,
		queue: make([]time.Time, 0, limit),
	})
	return len(m.rules) - 1
}

// AllowRequest checks if a new request can be served based on all active rules.
// Returns true only if the request passes all active rules.
func (m *MultiRuleRateLimiter) AllowRequest() bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()

	// 1. Cleanup expired requests and check limits for all rules
	for _, rule := range m.rules {
		cutOff := now.Add(-rule.ttl)

		// Find the index of the first non-expired request
		// Timestamps are strictly increasing, so we just need to find the first one > cutOff
		head := 0
		for head < len(rule.queue) && !rule.queue[head].After(cutOff) {
			head++
		}

		// Remove expired requests
		if head > 0 {
			// If all expired
			if head >= len(rule.queue) {
				rule.queue = rule.queue[:0]
			} else {
				// Shift remaining elements to the front
				// Using copy is safer to avoid memory leaks if we were slicing a large array,
				// but here we are modifying the slice in place.
				copy(rule.queue, rule.queue[head:])
				rule.queue = rule.queue[:len(rule.queue)-head]
			}
		}

		// Check if adding one more would exceed the limit
		if len(rule.queue) >= rule.limit {
			return false
		}
	}

	// 2. If all rules pass, record the request in all rules
	for _, rule := range m.rules {
		rule.queue = append(rule.queue, now)
	}

	return true
}
