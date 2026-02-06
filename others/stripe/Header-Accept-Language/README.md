# Accept-Language Header Parsing

Implement a parser for the HTTP `Accept-Language` header. Given a header
string and a list of supported languages, return the languages to serve in
priority order.

Rules:
- Each token may include a `q` factor (priority). Higher `q` comes first.
- A token without `q` has priority `1.0`.
- The wildcard `*` matches all remaining supported languages.
- A token may partially match a language prefix (e.g., `fr` matches `fr-CA`).
- Tokens with `q=0` exclude matching languages.

## Input

- `accept_language`: header value string.
- `supported_languages`: list of supported language tags.

## Output

- Ordered list of supported languages to serve.

## Example

**Input**
```
accept_language = "fr-FR;q=1, fr;q=0.5, *;q=0"
supported_languages = ["fr-FR", "fr-CA", "en-US"]
```

**Output**
```
["fr-FR", "fr-CA"]
```

## Constraints

- Supported languages are unique.
- `q` is in `[0, 1]`.
