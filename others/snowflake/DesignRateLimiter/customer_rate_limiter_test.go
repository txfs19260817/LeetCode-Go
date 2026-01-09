package snowflake

import (
	"testing"
	"time"
)

func TestCustomerRateLimiter_Example(t *testing.T) {
	// Initialize with ttl = 1000 ms and limit = 3.
	rl := NewCustomerRateLimiter(1000, 3)

	check := func(id int, expected bool, desc string) {
		if got := rl.AllowRequest(id); got != expected {
			t.Errorf("%s: Customer %d expected %v, got %v", desc, id, expected, got)
		}
	}

	// t = 0ms
	check(123, true, "t=0ms")
	check(456, true, "t=0ms")

	// t = 100ms
	time.Sleep(100 * time.Millisecond)
	check(123, true, "t=100ms")
	check(456, true, "t=100ms")

	// t = 200ms
	time.Sleep(100 * time.Millisecond)
	check(123, true, "t=200ms")
	check(456, true, "t=200ms")

	// t = 300ms
	// Limit of 3 is hit for both customers (requests at 0, 100, 200 are active)
	time.Sleep(100 * time.Millisecond)
	check(123, false, "t=300ms")
	check(456, false, "t=300ms")

	// t = 400ms
	time.Sleep(100 * time.Millisecond)
	check(123, false, "t=400ms")
	check(456, false, "t=400ms")

	// t = 500ms
	time.Sleep(100 * time.Millisecond)
	check(123, false, "t=500ms")
	check(456, false, "t=500ms")

	// t = 600ms
	time.Sleep(100 * time.Millisecond)
	check(123, false, "t=600ms")
	check(456, false, "t=600ms")

	// t = 700ms
	time.Sleep(100 * time.Millisecond)
	check(123, false, "t=700ms")
	check(456, false, "t=700ms")
}
