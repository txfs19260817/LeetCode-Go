# Search History System

- Difficulty: Hard
- Companies: Google, Microsoft
- Stages: Onsite
- Asked By: Google, Microsoft
- Source: https://www.hack2hire.com/companies/google/coding-questions/678878d0402dc47b67a14945/practice?questionId=678879df402dc47b67a14946

## Problem

When a user opens the search bar, the system should provide the most recent k search inputs as keyword recommendations. Each new search input should also be recorded and used for future recommendations. If a keyword is searched multiple times, only the most recent occurrence should be considered.

Implement the SearchHistory class:

SearchHistory(int capacity): Initialize the search history with a positive size capacity.
void add(String term): Add a search history to the cache. If the number of keys exceeds the capacity of this operation, evict the least recently used key.
List<String> getTopK(int k): Return the most recent k search inputs.

Constraints:

1 ≤ k ≤ 104
0 ≤ Number of search operations ≤ 105
Search keywords consist of lowercase and uppercase English letters.

Example:

Input:
["SearchHistory", "add", "add", "getTopK", "add", "getTopK","add","getTopK"]
[[3], ["paris"], ["tokyo"], [2], ["beijing"], [3], ["toronto"], [3]]
Output:
[null, null, null, ["tokyo","paris"], null, ["beijing", "tokyo", "paris"], null, ["toronto", "beijing", "tokyo"]
Explanation:

SearchHistory searchHistory = new SearchHistory(3); // Initialize with capacity 3
searchHistory.add("paris");
searchHistory.add("tokyo");
searchHistory.getTopK(2); // Returns ["tokyo,"paris"]
searchHistory.add("beijing");
searchHistory.getTopK(3); // Returns ["beijing","tokyo","paris"]
searchHistory.add("toronto");
searchHistory.getTopK(3); // Returns ["toronto","beijing","tokyo"]. Since the capacity is 3, "pairs" was evicted.

## Python Template

```python
from typing import List, Optional

class SearchHistory:
    def __init__(self, capacity: int):
        # TODO: Initialize SearchHistory

    def add(self, term: str) -> None:
        # TODO: Implement add logic

    def getTopK(self, k: int) -> List[str]:
        # TODO: Implement getTopK logic
```
