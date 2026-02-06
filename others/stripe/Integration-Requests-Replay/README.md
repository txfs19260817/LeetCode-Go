# Integration Requests Replay

You are given a list of HTTP request records captured from an integration.
Each record includes the URL path, headers, method, and body. Reconstruct and
replay each request against a base URL, then parse the JSON response.

## Input

- `requests`: list of request objects:
  - `url`, `headers`, `method`, `body`.
- `base_url`: base server URL.

## Output

- For each request, print or return the status code and parsed JSON response.

## Example

**Request**
```
{
  "url": "/v1/charges",
  "method": "POST",
  "body": "amount=123&currency=usd"
}
```

**Output**
```
status = 200
body = { ... }
```

## Constraints

- Methods are `GET` or `POST`.
- Body is `application/x-www-form-urlencoded`.
