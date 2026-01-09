package snowflake

import (
	"sync"
)

// CustomerRateLimiter implements a per-customer rate limiter.
type CustomerRateLimiter struct {
	ttl   int
	limit int
	// customers maps customerId to their specific RateLimiter using sync.Map for concurrent safety
	customers sync.Map
}

// NewCustomerRateLimiter initializes the rate limiter with a time-to-live window ttl in milliseconds
// and a maximum limit for the number of requests allowed per customer in that time frame.
func NewCustomerRateLimiter(ttl, limit int) *CustomerRateLimiter {
	return &CustomerRateLimiter{ttl: ttl, limit: limit}
}

// AllowRequest checks if a new request for the given customerId can be served
// based on recent requests for the same customer within the ttl window.
func (c *CustomerRateLimiter) AllowRequest(customerId int) bool {
	newRl := NewRateLimiter(c.ttl, c.limit)
	actualRl, _ := c.customers.LoadOrStore(customerId, newRl)

	return actualRl.(*RateLimiter).AllowRequest()
}
