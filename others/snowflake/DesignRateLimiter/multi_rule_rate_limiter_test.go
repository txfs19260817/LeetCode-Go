package snowflake

import (
	"testing"
	"time"
)

func TestMultiRuleRateLimiter_Example(t *testing.T) {
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
	// Rule 0 (500ms, 2) -> Active: [0, 100]. Limit reached.
	// Rule 1 (1000ms, 3) -> Active: [0, 100]. Limit ok.
	// Rejected by Rule 0.
	time.Sleep(100 * time.Millisecond)
	check("t=200ms", false)

	// t = 300ms
	// Rule 0 full.
	time.Sleep(100 * time.Millisecond)
	check("t=300ms", false)

	// t = 400ms
	// Rule 0 full.
	time.Sleep(100 * time.Millisecond)
	check("t=400ms", false)

	// t = 500ms
	// Rule 0: Cutoff 0. t=0 expires (0 <= 0). Active: [100]. Size 1. Limit 2. OK.
	// Rule 1: Cutoff -500. Active: [0, 100]. Size 2. Limit 3. OK.
	// Allowed.
	time.Sleep(100 * time.Millisecond)
	check("t=500ms", true)
	// After update: R0=[100, 500], R1=[0, 100, 500]

	// t = 600ms
	// Rule 0: Cutoff 100. t=100 expires (100 <= 100). Active: [500]. OK.
	// Rule 1: Cutoff -400. Active: [0, 100, 500]. Size 3. Limit 3. Fail.
	time.Sleep(100 * time.Millisecond)
	check("t=600ms", false)

	// t = 700ms
	time.Sleep(100 * time.Millisecond)
	check("t=700ms", false)

	// t = 800ms
	time.Sleep(100 * time.Millisecond)
	check("t=800ms", false)

	// t = 900ms
	time.Sleep(100 * time.Millisecond)
	check("t=900ms", false)

	// t = 1000ms
	// Rule 0: Cutoff 500. t=500 expires (500 <= 500). Active: []. OK.
	// Rule 1: Cutoff 0. t=0 expires (0 <= 0). Active: [100, 500]. Size 2. Limit 3. OK.
	// Allowed.
	time.Sleep(100 * time.Millisecond)
	check("t=1000ms", true)
}
