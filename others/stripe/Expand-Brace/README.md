# Expand Brace Expressions

Given a string that may contain brace expressions like `{a,b,c}`, expand it
into all possible strings by substituting each token.

If braces are invalid (missing, empty, or reversed), return the original
string unchanged. Multiple brace groups may appear sequentially.

## Input

- `input_str`: a string that may contain comma-separated tokens inside `{}`.

## Output

- A list of expanded strings.

## Example

**Input**
```
/{2021,2022}/{jan,feb}/report
```

**Output**
```
[
  "/2021/jan/report",
  "/2021/feb/report",
  "/2022/jan/report",
  "/2022/feb/report"
]
```

## Constraints

- Tokens contain no braces.
- Commas only separate tokens inside braces.
