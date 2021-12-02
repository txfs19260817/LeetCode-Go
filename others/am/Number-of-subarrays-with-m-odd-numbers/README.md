# Number of subarrays with m odd numbers

Given an array of `n` elements and an integer `m`, we need to write a program to find the number of contiguous subarrays
in the array which contains exactly `m` odd numbers.

```javascript
Input : arr = {2, 5, 6, 9}, m = 2
Output : 2
Explanation:
    subarrays
are [2, 5, 6, 9]
and [5, 6, 9]
```

```javascript
Input : arr = {2, 2, 5, 6, 9, 2, 11}, m = 2
Output : 8
Explanation:
    subarrays
are [2, 2, 5, 6, 9],
    [2, 5, 6, 9], [5, 6, 9], [2, 2, 5, 6, 9, 2],
    [2, 5, 6, 9, 2], [5, 6, 9, 2], [6, 9, 2, 11]
and [9, 2, 11]
```

https://www.geeksforgeeks.org/number-subarrays-m-odd-numbers/