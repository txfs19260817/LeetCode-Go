def fib():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b


def take(gen, n):
    return [next(gen) for _ in range(n)]


if __name__ == "__main__":
    g = fib()
    assert take(g, 7) == [0, 1, 1, 2, 3, 5, 8]
    assert take(g, 5) == [13, 21, 34, 55, 89]

    g2 = fib()
    assert take(g2, 0) == []
