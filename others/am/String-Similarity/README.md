# Sum of similarities of string with all of its suffixes

Given a string `str`, the task is to find the sum of the similarities of `str` with each of its suffixes. The similarity of
strings A and B is the length of the longest prefix common to both the strings i.e. the similarity of “aabc” and “aab”
is 3 and that of “qwer” and “abc” is 0.

```javascript
Input: str = 'ababa'
Output: 9
```

The suffixes of str are “ababa”, “baba”, “aba”, “ba” and “a”. The similarities of these strings with the original string
“ababa” are 5, 0, 3, 0 & 1 respectively. Thus, the answer is 5 + 0 + 3 + 0 + 1 = 9.

```javascript
Input: str = 'aaabaab' 
Output: 13 
```

https://www.geeksforgeeks.org/sum-of-similarities-of-string-with-all-of-its-suffixes/
https://www.hackerrank.com/challenges/string-similarity/topics