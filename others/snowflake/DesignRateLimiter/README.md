# Design Rate Limiter

## Design Rate Limiter

- **Difficulty:** Medium
- **Tags:** Design, Queue
- **Interview Stages:** Onsite, Screening
- **Frequency:** 70%
- **Asked By:** Microsoft, Atlassian
- **Last Reported:** 3 months ago

Design a rate limiter system that begins with a specified capacity (`limit`) indicating the maximum number of requests
allowed within a certain time interval (`ttl`, unit: Millisecond). When a new request arrives, return `true` if it is
accepted under the current limit. Otherwise, return `false`.

Implement the `RateLimiter` class:

- `RateLimiter(int ttl, int limit)`  
  Initializes the rate limiter with a time-to-live window `ttl` in milliseconds and a maximum `limit` for the number of
  requests allowed in that time frame.

- `boolean allowRequest()`  
  Checks if a new request can be served based on recent requests within the `ttl` window.

### Example 1

> **Input:**  
> *(Assuming the interval for each request is **100ms**)*
>
>
`["RateLimiter","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest"]`  
> `[[1000,5],[],[],[],[],[],[],[],[],[],[],[]]`
>
> **Output:**  
> `[null, true, true, true, true, true, false, false, false, false, false, true]`
>
> **Explanation:**
> - `RateLimiter rateLimiter = new RateLimiter(1000, 5);` // Initialize with ttl = 1000 ms and limit = 5.
> - `rateLimiter.allowRequest();` // Returns true. (t = 0ms)
> - `rateLimiter.allowRequest();` // Returns true. (t = 100ms)
> - `rateLimiter.allowRequest();` // Returns true. (t = 200ms)
> - `rateLimiter.allowRequest();` // Returns true. (t = 300ms)
> - `rateLimiter.allowRequest();` // Returns true. (t = 400ms)
> - `rateLimiter.allowRequest();` // Returns false, as the limit of 5 is hit within the last 1000 ms. (t = 500ms)
> - `rateLimiter.allowRequest();` // Returns false. (t = 600ms)
> - `rateLimiter.allowRequest();` // Returns false. (t = 700ms)
> - `rateLimiter.allowRequest();` // Returns false. (t = 800ms)
> - `rateLimiter.allowRequest();` // Returns false. (t = 900ms)
> - `rateLimiter.allowRequest();` // Returns true, as the oldest request (t = 0ms) has expired. (t = 1000ms)

### Example 2

> **Input:**  
> *(Assuming the interval for each request is **300ms**)*
>
>
`["RateLimiter","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest","allowRequest"]`  
> `[[1000,5],[],[],[],[],[],[],[],[],[],[],[]]`
>
> **Output:**  
> `[null, true, true, true, true, true, true, true, true, true, true]`
>
> **Explanation:**  
> Since all 10 requests were sent at a rate slower than the rate limiter, they should all be allowed.

---

## Design Rate Limiter II

**Difficulty:** Hard  
**Topics:** Design, Queue

---

## Metadata

- **Interview Stages:** Onsite  
- **Frequency:** ~80%  
- **Asked By:** Snowflake  
- **Last Reported:** 6 days ago  

---

## Problem Description

*(This question is a variation of the Hack2Hire question **Design Rate Limiter**. If you haven't completed that question yet, it is recommended to solve it first.)*

Design a rate limiter system that controls the number of requests allowed within specified time windows. The system should support multiple rate-limiting rules simultaneously and determine if a new request should be accepted or rejected based on all active rules.

A rate-limiting rule consists of:

- A time window (`ttl`) in milliseconds
- A maximum number of requests (`limit`) allowed within that window

---

## API Design

Implement the `MultiRuleRateLimiter` class:

- `MultiRuleRateLimiter()`  
  Initializes an empty rate limiter with no rules.

- `int registerRule(int ttl, int limit)`  
  Adds a new rate-limiting rule with the specified `ttl` and `limit`.  
  Returns the rule ID (index) for future reference.

- `boolean allowRequest()`  
  Checks if a new request can be served based on all active rules.  
  Returns `true` only if the request passes all active rules.  
  If no rules are registered, returns `true`.

---

## Constraints

- `1 ≤ ttl ≤ 10^9` milliseconds
- `1 ≤ limit ≤ 10^6` requests
- The system must be thread-safe
- Assume the rule ID starts from `0`

---

## Example

**Input:**  
*(Assuming the interval for each request is 100ms)*

```text
["MultiRuleRateLimiter", "registerRule", "registerRule", "allowRequest",
 "allowRequest", "allowRequest", "allowRequest", "allowRequest",
 "allowRequest", "allowRequest", "allowRequest", "allowRequest",
 "allowRequest", "allowRequest"]

[[], [500, 2], [1000, 3], [], [], [], [], [], [], [], [], [], [], []]
````

**Output:**

```text
[null, 0, 1, true, true, false, false, false, true,
 false, false, false, false, true]
```

**Explanation:**

* `MultiRuleRateLimiter limiter = new MultiRuleRateLimiter();`
* `limiter.registerRule(500, 2);`
  Returns `0`. Registers rule 0: allow 2 requests per 500ms.
* `limiter.registerRule(1000, 3);`
  Returns `1`. Registers rule 1: allow 3 requests per 1000ms.
* `limiter.allowRequest();` → `true` (t = 0ms)
* `limiter.allowRequest();` → `true` (t = 100ms)
* `limiter.allowRequest();` → `false`, rejected by rule 0 (t = 200ms)
* `limiter.allowRequest();` → `false`, rejected by rule 0 (t = 300ms)
* `limiter.allowRequest();` → `false`, rejected by rule 0 (t = 400ms)
* `limiter.allowRequest();` → `true` (t = 500ms)
* `limiter.allowRequest();` → `false`, rejected by rule 1 (t = 600ms)
* `limiter.allowRequest();` → `false`, rejected by rule 1 (t = 700ms)
* `limiter.allowRequest();` → `false`, rejected by rule 1 (t = 800ms)
* `limiter.allowRequest();` → `false`, rejected by rule 1 (t = 900ms)
* `limiter.allowRequest();` → `true` (t = 1000ms)

