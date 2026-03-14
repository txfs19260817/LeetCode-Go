# Date Validator+1

- Difficulty: Easy
- Companies: Google
- Stages: Onsite
- Asked By: Google
- Source: https://www.hack2hire.com/companies/google/coding-questions/6764929acc8f5deed53ccd75/practice?questionId=6764932ccc8f5deed53ccd76

## Problem

Date Validator
+1
Easy
String
Interview Stages
Onsite
Frequency
Asked By
Last Reported
3 weeks ago

Given three integers, determine whether they can form a valid date in any standard date format. A date is considered valid if the integers can represent a date in formats such as yyyy-MM-dd, MM-dd-yyyy, or dd-MM-yyyy. All possible valid date formats should be considered, and the year should be within the range 1 to 9999. The month must be between 1 and 12, and the day must be valid for the given month and year, taking leap years into account.

Constraints:

1 <= a, b, c <= 9999

Example 1:

Input: a = 2020, b = 2, c = 29
Output: true
Explanation: The integers can form the date February 29, 2020 (2020-02-29), which is a valid leap day.

Example 2:

Input: a = 2019, b = 2, c = 29
Output: false
Explanation: 2019 is not a leap year, so February 29, 2019 is invalid.

Example 3:

Input: a = 12, b = 31, c = 2021
Output: true
Explanation: The integers can form December 31, 2021 (2021-12-31), which is a valid date.

## Python Template

```python
from typing import List, Optional

class Solution:
    def isValidDate(self, a: int, b: int, c: int) -> bool:
        # TODO: Implement isValidDate logic
```
