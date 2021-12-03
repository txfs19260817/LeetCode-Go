# Find 0's rectangles

https://leetcode.com/discuss/interview-question/1062462/indeed-karat-questions

Give a matrix of 1's and 0's. For example

```javascript
matrix = [
    ["1", "1", "1", "1", "1"],
    ["1", "1", "0", "0", "1"],
    ["1", "1", "0", "0", "1"],
    ["1", "1", "1", "1", "1"]]
```

Round 1: find the rectangle that is made of 0s, either return the start and end index OR height and length of the
rectangle. There is only 1 rectangle in each matrix. I solved it by looping to find the first and last zero.

Round 2: same problem but now the matrix may contains many rectangles. Return the start and end indexes of each rec in
an array.

```javascript
matrix = [
    ["0", "1", "1", "1", "1"],
    ["1", "1", "0", "0", "1"],
    ["0", "1", "0", "0", "1"],
    ["0", "1", "1", "1", "1"],
    ["1", "0", "1", "1", "1"]] 
```

4 rectangles in here . 0 by itself is also a rectangle, and vertical rectangle also counts.

