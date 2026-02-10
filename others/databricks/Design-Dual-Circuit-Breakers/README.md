# Design Dual Circuit Breakers

Simulate a distributed system gateway with primary and secondary servers, each with its own circuit breaker.

## Concepts

**Circuit breaker states:** Closed (requests pass through) or Open (requests blocked).

**Server:** has a fixed list of boolean outcomes. `handle(requestId)` returns `outcomes[requestId]`.

**CircuitBreaker(server, failureThreshold, resetThreshold):**
- **Closed:** requests go to server. On failure, increment consecutive failures. When `failures >= failureThreshold`, circuit opens (reset `rejectedCount` to 0).
- **Open:** skip server, increment `rejectedCount`. When `rejectedCount >= resetThreshold`, circuit closes (reset failures to 0). The request is NOT retried on this reset.

**Gateway(primaryBreaker, secondaryBreaker):**
`routeRequests(totalRequests)` — process requests 0 to totalRequests-1.

## Routing Logic

1. If primary is **closed** → attempt primary server
   - Success → reset primary failures. Done. Result: `"Primary"`
   - Fail → increment failures (open if threshold). Try secondary.
2. If primary is **open** → increment rejectedCount (close if threshold). Try secondary.
3. If secondary is **closed** → attempt secondary server, update state.
4. If secondary is **open** → increment rejectedCount (close if threshold).

## Result String

- Both attempted: `"Primary -> Secondary"`
- Only primary: `"Primary"`
- Only secondary: `"Secondary"`
- Neither: `"Rejected"`

("Attempted" means `server.handle()` was called.)

## Example

**Input:**
```
primary outcomes   = [true, false, false, true, true, false, true]
secondary outcomes = [false, true, false, false, true, true, true]
failureThreshold = 2, resetThreshold = 2 (both breakers)
totalRequests = 7
```

**Output:**
```
["Primary", "Primary -> Secondary", "Primary -> Secondary", "Secondary", "Rejected", "Primary", "Primary"]
```

**Trace:**
| Req | Primary State | Action | Secondary State | Action | Result |
|-----|---------------|--------|-----------------|--------|--------|
| 0 | Closed | handle=T, fail=0 | — | — | Primary |
| 1 | Closed | handle=F, fail=1 | Closed | handle=T, fail=0 | Primary -> Secondary |
| 2 | Closed | handle=F, fail=2→Open | Closed | handle=F, fail=1 | Primary -> Secondary |
| 3 | Open | rej=1 | Closed | handle=F, fail=2→Open | Secondary |
| 4 | Open | rej=2→Close | Open | rej=1 | Rejected |
| 5 | Closed | handle=F, fail=1 | Open | rej=2→Close | Primary |
| 6 | Closed | handle=T, fail=0 | — | — | Primary |
