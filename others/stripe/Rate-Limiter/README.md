# Rate Limiter

Design a per-user rate limiter that allows at most `max_requests` within a
rolling time window of `window_seconds`.

Provide:

- `is_allowed(user_id)` to check and record a request.
- `get_remaining_requests(user_id)` to return remaining quota and time until
  reset.

## Input

- `max_requests`: maximum allowed requests in window.
- `window_seconds`: window size in seconds.
- `user_id`: unique user identifier.

## Output

- Boolean allow/deny per request.
- Remaining count and seconds until reset.

## Example

With `max_requests=5` and `window_seconds=2`, the 6th request in a 2-second
interval is rejected.

## Constraints

- Time is based on current wall-clock time.
- User IDs are strings.
