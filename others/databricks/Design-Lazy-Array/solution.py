class LazyArray:
    def __init__(self, arr, fns=None):
        self._arr = arr
        self._fns = fns if fns is not None else []

    def map(self, fn):
        """Return a new LazyArray with fn appended to the pipeline."""
        return LazyArray(self._arr, self._fns + [fn])

    def indexOf(self, target):
        """Apply all accumulated functions and return index of target, or -1."""
        for i, val in enumerate(self._arr):
            result = val
            for fn in self._fns:
                result = fn(result)
            if result == target:
                return i
        return -1


if __name__ == "__main__":
    # Example 1: single map, indexOf
    la = LazyArray([10, 20, 30, 40, 50])
    doubled = la.map(lambda n: n * 2)
    assert doubled.indexOf(40) == 1

    # Example 2: chained maps
    tripled = doubled.map(lambda n: n * 3)
    assert tripled.indexOf(240) == 3

    # Example 3: not found
    la2 = LazyArray([1, 2, 3, 4, 5])
    mapped = la2.map(lambda n: n + 10)
    assert mapped.indexOf(100) == -1

    # Chain of 3 maps
    la3 = LazyArray([1, 2, 3, 4, 5])
    result = la3.map(lambda n: n + 1).map(lambda n: n * 3).map(lambda n: n - 1)
    assert result.indexOf(5) == 0
    assert result.indexOf(11) == 2
    assert result.indexOf(17) == 4
    assert result.indexOf(99) == -1

    # Single element
    la4 = LazyArray([42])
    m4 = la4.map(lambda n: n * 2)
    assert m4.indexOf(84) == 0
    assert m4.indexOf(42) == -1

    # Original unchanged after map
    la5 = LazyArray([10, 20, 30])
    _ = la5.map(lambda n: n * 100)
    assert la5.indexOf(20) == 1

    print("All tests passed!")
