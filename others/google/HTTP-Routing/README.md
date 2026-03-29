# HTTP Routing

Implement a small HTTP router with two public functions:

1. `registerHandler(path)` registers a route pattern.
2. `getHandler(path)` returns the matching handler identifier or `None`.

Routes are registered using path patterns such as:

```text
/home
/users/new
/users/{id}
/users/{id}/pics/{picId}
```

Matching rules:

- A literal segment must match exactly.
- A parameter segment such as `{id}` matches any single non-empty segment.
- Exact literal routes take priority over parameter routes.
- A trailing slash that creates an empty segment is invalid for lookup, so `/users/` returns `None`.

Handler identifier convention used by this implementation:

- A single static segment route like `/home` maps to `"home"`.
- Any route containing multiple segments or parameters maps to its normalized pattern string.

## Example
**Input:**
```text
registerHandler("/home")
registerHandler("/users/new")
registerHandler("/users/{id}")
registerHandler("/users/{id}/pics/{picId}")

getHandler("/home")
getHandler("/users/101")
getHandler("/users/")
```

**Output:**
```text
"home"
"/users/{id}"
None
```
