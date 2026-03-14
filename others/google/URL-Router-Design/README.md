# URL Router Design

- Difficulty: Hard
- Companies: Google, Atlassian
- Stages: Onsite
- Asked By: Google, Atlassian
- Source: https://www.hack2hire.com/companies/google/coding-questions/68fa76eb0fa32e4d82f4fc1a/practice?questionId=68fa7c5e0fa32e4d82f4fc1f

## Problem

Design a URL router that maps handler function names to registered URL patterns. The router must support exact path matching and also allow using the '*' wildcard in any segment of a pattern, where '*' matches any single segment in a query path.

When handling a query, the router should return the handler function name corresponding to the most specific matching pattern:

Prefer the match with the fewest wildcards.
If no pattern matches, return the empty string "".

Implement the UrlRouter class:

UrlRouter() Initializes an empty URL router.

void put(String pattern, String funcName) Registers the handler function name funcName for the pattern pattern.

Patterns may use the '*' wildcard in any segment. A pattern can have any number of wildcards.
If the same pattern is registered again, its handler is updated to the latest funcName.

String get(String url) Returns the handler function name for the most specific registered pattern matching the input url.

If multiple patterns match, return the one with the fewest wildcards.
If no pattern matches, return the empty string "".

Constraints:

All url and pattern values start with '/' and consist of one or more segments separated by '/'.
funcName is a non-empty string with at most 20 lowercase English letters, and different patterns can use the same funcName.
1 ≤ total calls to put and get ≤ 104.

Example

Input:
["UrlRouter", "put", "put", "put", "put", "get", "get", "get", "get", "get", "get"]

[[], ["/*", "oneWildcard"], ["/abc/*", "abcOneWildcard"], ["/abc/*/*", "abcTwoWildcards"], ["/abc/bcd", "exactMatch"], ["/abc/bcd"], ["/abc/def"], ["/xyz"], ["/abc/def/ghi"], ["/def/ghi"], ["/abc"]]

Output:
[null, null, null, null, null, "exactMatch", "abcOneWildcard", "oneWildcard", "abcTwoWildcards", "", "oneWildcard"]

Explanation:

UrlRouter router = new UrlRouter();
router.put("/*", "oneWildcard");
router.put("/abc/*", "abcOneWildcard");
router.put("/abc/*/*", "abcTwoWildcards");
router.put("/abc/bcd", "exactMatch");
router.get("/abc/bcd"); // Returns "exactMatch". Exact match takes priority over wildcards.
router.get("/abc/def"); // Returns "abcOneWildcard". Matches both "'/*" and "/abc/*", but "/abc/*" is more specific (same wildcard count, but longer prefix).
router.get("/xyz"); // Returns "oneWildcard". Only "/*" matches.
router.get("/abc/def/ghi"); // Returns "abcTwoWildcards". Only "/abc/*/*" matches this three-segment path.
router.get("/def/ghi"); // Returns "".
router.get("/abc"); // Returns "oneWildcard". Only "/*" matches.

## Python Template

```python
from typing import List, Optional

class UrlRouter:
    def __init__(self, ):
        # TODO: Initialize UrlRouter

    def put(self, pattern: str, funcName: str) -> None:
        # TODO: Implement put logic

    def get(self, url: str) -> str:
        # TODO: Implement get logic
```
