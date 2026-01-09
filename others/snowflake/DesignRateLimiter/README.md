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

