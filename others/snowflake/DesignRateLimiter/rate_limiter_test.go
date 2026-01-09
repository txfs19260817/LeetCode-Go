package snowflake

import (
	"testing"
	"time"
)

func TestRateLimiter_Example1(t *testing.T) {
	// Initialize with ttl = 1000 ms and limit = 5.
	rl := NewRateLimiter(1000, 5)

	// Step 1: t = 0ms
	if !rl.AllowRequest() {
		t.Errorf("Request at t=0ms should be allowed")
	}

	// Step 2: t = 100ms
	time.Sleep(100 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request at t=100ms should be allowed")
	}

	// Step 3: t = 200ms
	time.Sleep(100 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request at t=200ms should be allowed")
	}

	// Step 4: t = 300ms
	time.Sleep(100 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request at t=300ms should be allowed")
	}

	// Step 5: t = 400ms
	time.Sleep(100 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request at t=400ms should be allowed")
	}

	// Step 6: t = 500ms
	// Limit of 5 is hit (0, 100, 200, 300, 400 are active)
	time.Sleep(100 * time.Millisecond)
	if rl.AllowRequest() {
		t.Errorf("Request at t=500ms should be denied")
	}

	// Step 7: t = 600ms
	time.Sleep(100 * time.Millisecond)
	if rl.AllowRequest() {
		t.Errorf("Request at t=600ms should be denied")
	}

	// Step 8: t = 700ms
	time.Sleep(100 * time.Millisecond)
	if rl.AllowRequest() {
		t.Errorf("Request at t=700ms should be denied")
	}

	// Step 9: t = 800ms
	time.Sleep(100 * time.Millisecond)
	if rl.AllowRequest() {
		t.Errorf("Request at t=800ms should be denied")
	}

	// Step 10: t = 900ms
	time.Sleep(100 * time.Millisecond)
	if rl.AllowRequest() {
		t.Errorf("Request at t=900ms should be denied")
	}

	// Step 11: t = 1000ms
	// Oldest request (t=0ms) expires at t=1000ms (cutoff = 0)
	time.Sleep(100 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request at t=1000ms should be allowed")
	}
}

func TestRateLimiter_Example2(t *testing.T) {
	// Initialize with ttl = 1000 ms and limit = 5.
	rl := NewRateLimiter(1000, 5)

	// We have 10 requests total, interval 300ms.
	// t=0, 300, 600, 900, 1200, 1500, 1800, 2100, 2400, 2700

	// 1. t=0 allowed
	if !rl.AllowRequest() {
		t.Errorf("Request 1 at t=0ms should be allowed")
	}

	// 2. t=300 allowed
	time.Sleep(300 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request 2 at t=300ms should be allowed")
	}

	// 3. t=600 allowed
	time.Sleep(300 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request 3 at t=600ms should be allowed")
	}

	// 4. t=900 allowed
	time.Sleep(300 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request 4 at t=900ms should be allowed")
	}

	// 5. t=1200 allowed
	// At t=1200, cutoff is 200.
	// t=0 is expired.
	// Active: 300, 600, 900. (3 items)
	// New item 1200 added. Total 4. Allowed.
	time.Sleep(300 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request 5 at t=1200ms should be allowed")
	}

	// 6. t=1500 allowed
	// At t=1500, cutoff is 500.
	// t=300 is expired.
	// Active: 600, 900, 1200. (3 items)
	// New item 1500 added. Total 4. Allowed.
	time.Sleep(300 * time.Millisecond)
	if !rl.AllowRequest() {
		t.Errorf("Request 6 at t=1500ms should be allowed")
	}

	// ... and so on. All should be allowed because we drop one every step after reaching capacity?
	// Wait. Capacity is 5.
	// t=0: [0]
	// t=300: [0, 300]
	// t=600: [0, 300, 600]
	// t=900: [0, 300, 600, 900]
	// t=1200: Drop 0 (<=200). [300, 600, 900]. Add 1200. -> [300, 600, 900, 1200] (4 items)
	// t=1500: Drop 300 (<=500). [600, 900, 1200]. Add 1500. -> [600, 900, 1200, 1500] (4 items)
	// It seems we stay at 4 items constant after warmup.
	// So all 10 should be allowed.

	for i := 0; i < 4; i++ { // remaining 4 requests to make 10 total
		time.Sleep(300 * time.Millisecond)
		if !rl.AllowRequest() {
			t.Errorf("Request %d should be allowed", i+7)
		}
	}
}
