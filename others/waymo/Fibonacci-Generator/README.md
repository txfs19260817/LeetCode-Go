# Fibonacci generator

Implement a Fibonacci generator with:
- lazy generation
- O(1) extra space

要求：惰性生成、O(1) 额外空间。

Python interface:
```python
def fib():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b
```

## Example
**Input:**
```
Take first 7 values from fib()
```

**Output:**
```
[0, 1, 1, 2, 3, 5, 8]
```
