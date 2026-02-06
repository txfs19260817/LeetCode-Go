# Load Balancer With Weights

Implement a load balancer that routes each request to the server with the
lowest current load. Each request has a `weight` and an optional `ttl`
indicating how long it contributes to server load.

When a request is routed, add its weight to the server. When the TTL expires,
that weight is removed from the server.

## Input

- `servers`: list of server IDs.
- `weight`: integer weight for the request.
- `ttl`: time-to-live in seconds.

## Output

- The server ID chosen for each request.

## Example

**Input**
```
servers = ["s1", "s2", "s3"]
requests = [(5, 10), (3, 5), (2, 20)]
```

**Output**
```
["s1", "s2", "s3"]
```

## Constraints

- Weights are non-negative integers.
- TTLs are non-negative integers.
