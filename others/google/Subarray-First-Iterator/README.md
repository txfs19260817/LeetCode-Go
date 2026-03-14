# Subarray-First Iterator

- Difficulty: Easy
- Companies: Google
- Stages: Screening
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/677be9f75cbaff553b53b067/practice?questionId=677bf4df5cbaff553b53b068

## Problem

Given two integer arrays, a and b. Design a custom iterator that iterates through the elements in the following order:

First, yield all elements from array b in their original order.
Then, yield the remaining elements from array a that are not present in array b, maintaining their original order from a.

Note that

All elements from b are prioritized and appear first in the iteration.
The relative order of elements from both arrays is preserved.

Constraints:

1
≤
a.length
≤
10
5
1≤a.length≤10
5
1
≤
b.length
≤
𝑎
.
𝑙
𝑒
𝑛
𝑔
𝑡
ℎ
1≤b.length≤a.length
0
≤
a[i]
,
b[i]
≤
10
9
0≤a[i],b[i]≤10
9

Example 1:

Input: a = [1, 2, 3, 4, 5], b = [3, 5]
Output: [3, 5, 1, 2, 4]
Explanation: The iterator first outputs 3 and 5 from b, then yields the remaining elements of a [1, 2, 4].

Example 2:

Input: a = [10, 20, 30, 40], b = [20, 50]
Output: [20, 50, 10, 30, 40]

Example 3:

Input: a = [7, 8, 9], b = [1, 2, 3]
Output: [1, 2, 3, 7, 8, 9]

## Python Template

```python
class Solution:
    def __init__(self, a, b):
        # TODO: Implement __init__ logic.
        pass

    def hasNext(self):
        # TODO: Implement hasNext logic.
        pass

    def next(self):
        # TODO: Implement next logic.
        pass


if __name__ == '__main__':
    a1 = [1, 2, 3, 4, 5]
```
