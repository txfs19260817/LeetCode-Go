# Integration Request Runner

Reconstruct and execute HTTP requests from stored integration data.
Each request includes URL path, headers, method, and form-encoded body.
Return or print response status and JSON payload.

## Input

- `requests`: list of request records.
- `base_url`: base server URL.

## Output

- Status code and parsed JSON per request.

## Example

**Request**
```
{ "url": "/v1/charges", "method": "POST", "body": "amount=123&currency=usd" }
```

**Output**
```
status = 200
body = { ... }
```

## Constraints

- Methods are `GET` or `POST`.
- Body is URL-encoded key-value pairs.
