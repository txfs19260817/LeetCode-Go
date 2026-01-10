package snowflake

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMultiRuleRateLimiter_Example(t *testing.T) {
	t.Parallel()
	limiter := NewMultiRuleRateLimiter()

	// Register Rule 0: Allow 2 requests per 500ms.
	id0 := limiter.RegisterRule(500, 2)
	if id0 != 0 {
		t.Errorf("Expected rule ID 0, got %d", id0)
	}

	// Register Rule 1: Allow 3 requests per 1000ms.
	id1 := limiter.RegisterRule(1000, 3)
	if id1 != 1 {
		t.Errorf("Expected rule ID 1, got %d", id1)
	}

	check := func(desc string, expected bool) {
		if got := limiter.AllowRequest(); got != expected {
			t.Errorf("%s: expected %v, got %v", desc, expected, got)
		}
	}

	// t = 0ms
	check("t=0ms", true)

	// t = 100ms
	time.Sleep(100 * time.Millisecond)
	check("t=100ms", true)

	// t = 200ms
	// Rule 0 (500ms, 2) -> Active: [0, 100]. Limit reached (2).
	// Rule 1 (1000ms, 3) -> Active: [0, 100]. Limit ok (2 < 3).
	// Rejected by Rule 0.
	time.Sleep(100 * time.Millisecond)
	check("t=200ms", false)

	// t = 300ms
	// Rule 0 still full (oldest is at t=0, expires at t=500).
	time.Sleep(100 * time.Millisecond)
	check("t=300ms", false)

	// t = 400ms
	// Rule 0 still full.
	time.Sleep(100 * time.Millisecond)
	check("t=400ms", false)

	// t = 500ms
	// Rule 0: Cutoff 0. t=0 expires (0 <= 0). Active: [100]. Size 1. Limit 2. OK.
	// Rule 1: Cutoff -500. Active: [0, 100]. Size 2. Limit 3. OK.
	// Allowed.
	// Note: We need to be careful with exact boundary.
	// Sleep guarantees we are >= 500ms from start.
	time.Sleep(100 * time.Millisecond)
	check("t=500ms", true)

	// t = 600ms
	// Rule 0: Cutoff 100. t=100 expires (100 <= 100). Active: [500]. OK.
	// Rule 1: Cutoff -400. Active: [0, 100, 500]. Size 3. Limit 3. Fail.
	// Rejected by Rule 1.
	time.Sleep(100 * time.Millisecond)
	check("t=600ms", false)
}

func TestMultiRuleRateLimiter_ConcurrentAccess(t *testing.T) {
	t.Parallel()
	limiter := NewMultiRuleRateLimiter()
	// Allow 100 requests per 1 second
	limit := 100
	limiter.RegisterRule(1000, limit)

	var allowedCount int32
	var wg sync.WaitGroup
	totalRequests := 200 // Twice the limit

	// Start all goroutines effectively at the same time
	start := make(chan struct{})

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			if limiter.AllowRequest() {
				atomic.AddInt32(&allowedCount, 1)
			}
		}()
	}

	close(start)
	wg.Wait()

	// Since all requests happen nearly instantly, only 'limit' should pass.
	// We allow a small margin of error if time shifts during execution,
	// but with 1-second window, it should be exact unless execution takes > 1s.
	if int(allowedCount) != limit {
		t.Errorf("Concurrent access: expected %d allowed, got %d", limit, allowedCount)
	}
}

func TestMultiRuleRateLimiter_StrictOrder(t *testing.T) {
	t.Parallel()
	limiter := NewMultiRuleRateLimiter()
	// 5 requests per 200ms
	limiter.RegisterRule(200, 5)

	// Send 5 requests, all should pass
	for i := 0; i < 5; i++ {
		if !limiter.AllowRequest() {
			t.Errorf("Request %d should be allowed", i)
		}
	}

	// 6th request should fail immediately
	if limiter.AllowRequest() {
		t.Errorf("6th request should be denied immediately")
	}

	// Wait for expiration
	time.Sleep(210 * time.Millisecond)

	// Should be allowed again
	if !limiter.AllowRequest() {
		t.Errorf("Request after expiration should be allowed")
	}
}

func TestMultiRuleRateLimiter_MultipleRulesHierarchy(t *testing.T) {
	t.Parallel()
	limiter := NewMultiRuleRateLimiter()
	// Rule 1: 10 per 1 second (Long term)
	// Rule 2: 2 per 100 ms (Short burst)
	limiter.RegisterRule(1000, 10)
	limiter.RegisterRule(100, 2)

	// Burst 1: 2 requests
	if !limiter.AllowRequest() {
		t.Error("1st req failed")
	}
	if !limiter.AllowRequest() {
		t.Error("2nd req failed")
	}

	// 3rd request (immediate) -> Should fail due to Rule 2
	if limiter.AllowRequest() {
		t.Error("3rd req should fail due to burst limit")
	}

	// Wait 100ms
	time.Sleep(110 * time.Millisecond)

	// Burst 2: 2 requests
	if !limiter.AllowRequest() {
		t.Error("4th req failed after wait")
	}
	if !limiter.AllowRequest() {
		t.Error("5th req failed after wait")
	}

	// Now we have 4 requests total in < 1 sec.
	// Repeat until we hit long term limit (10).
	// We did 2 + 2 = 4.
	// Do 3 more bursts of 2.
	for i := 0; i < 3; i++ {
		time.Sleep(110 * time.Millisecond)
		if !limiter.AllowRequest() {
			t.Fatalf("Burst %d, req 1 failed", i+3)
		}
		if !limiter.AllowRequest() {
			t.Fatalf("Burst %d, req 2 failed", i+3)
		}
		if limiter.AllowRequest() {
			t.Fatalf("Burst %d, req 3 should fail", i+3)
		}
	}
	// Total requests accepted: 4 + 6 = 10.

	// Wait 100ms. Burst limit clears, but Long term limit (10) is full.
	time.Sleep(110 * time.Millisecond)

	if limiter.AllowRequest() {
		t.Error("Should be blocked by long term limit (10 per 1s)")
	}
}

func TestMultiRuleRateLimiter_SubTests(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		rules    [][2]int // pairs of {ttl, limit}
		requests []struct {
			delay   time.Duration
			allowed bool
		}
	}{
		{
			name:  "Single Rule Limit 1",
			rules: [][2]int{{100, 1}},
			requests: []struct {
				delay   time.Duration
				allowed bool
			}{
				{0, true},
				{0, false},
				{110 * time.Millisecond, true},
			},
		},
		{
			name:  "Overlapping Rules",
			rules: [][2]int{{100, 5}, {200, 2}}, // Rule 2 is stricter over longer period? No.
			// Rule 1: 5 per 100ms. Rule 2: 2 per 200ms.
			// Effectively capped at 2 per 200ms for sustained, but allowed 2 bursts?
			// Actually 2 per 200ms is stricter than 5 per 100ms (which is 10 per 200ms).
			// So Rule 2 dominates.
			requests: []struct {
				delay   time.Duration
				allowed bool
			}{
				{0, true},
				{0, true},
				{0, false},                      // Hit Rule 2 limit
				{110 * time.Millisecond, false}, // t=110. Rule 1 clears (expired). Rule 2 (200ms) not expired.
				// R2: [0, 0]. t=110. cutoff = 110-200 = -90. 0 > -90. Not expired.
				// So still full.
				{100 * time.Millisecond, true}, // t=210. R2 cutoff 10. 0 expired.
			},
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			limiter := NewMultiRuleRateLimiter()
			for _, r := range tc.rules {
				limiter.RegisterRule(r[0], r[1])
			}

			for i, req := range tc.requests {
				if req.delay > 0 {
					time.Sleep(req.delay)
				}
				if got := limiter.AllowRequest(); got != req.allowed {
					t.Errorf("Request %d: expected %v, got %v", i, req.allowed, got)
				}
			}
		})
	}
}
