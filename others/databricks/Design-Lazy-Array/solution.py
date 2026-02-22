from typing import Callable, Iterable


MapFn = Callable[[int], int]
FilterFn = Callable[[int], bool]
Op = tuple[str, MapFn | FilterFn]


class LazyArray:
    _DROPPED = object()

    def __init__(self, arr: list[int], ops: Iterable[Op] = ()):
        self._arr = arr
        self._ops: list[Op] = list(ops)

        # Incremental cache for this pipeline instance.
        self._evaluated_upto = -1
        self._cache: list[object] = []
        self._first_index: dict[int, int] = {}

    def map(self, fn: MapFn):
        """Return a new LazyArray with fn appended to the pipeline.

        Time: O(k), Space: O(k), where k is current pipeline length.
        """
        return LazyArray(self._arr, self._ops + [("map", fn)])

    def filter(self, predicate: FilterFn):
        """Return a new LazyArray with predicate appended to the pipeline.

        Time: O(k), Space: O(k), where k is current pipeline length.
        """
        return LazyArray(self._arr, self._ops + [("filter", predicate)])

    def indexOf(self, target: int) -> int:
        """Apply all accumulated functions and return index of target, or -1.

        Time: O(n*k) worst case, Space: O(1) extra,
        where n is array length and k is pipeline length.
        """
        for i, v in enumerate(self._arr):
            for t, func in self._ops:
                if t == "map":
                    v = func(v)
                elif t == "filter":
                    if not func(v):   # drop when predicate fails
                        v = None
                        break

            if v == target:  # check only after all ops for this item
                return i

        return -1

    def indexOf2(self, target: int) -> int:
        """Cached indexOf variant.

        First query: O(n*k) worst case.
        Repeated queries: only process the not-yet-cached suffix.
        Space: O(n) cache for this pipeline instance.
        """
        if target in self._first_index:
            return self._first_index[target]

        start = self._evaluated_upto + 1
        for i in range(start, len(self._arr)):
            v = self._arr[i]
            for t, func in self._ops:
                if t == "map":
                    v = func(v)
                elif t == "filter":
                    if not func(v):   # drop when predicate fails
                        v = self._DROPPED
                        break

            self._cache.append(v)
            self._evaluated_upto = i

            if v is self._DROPPED:
                continue

            if v not in self._first_index:
                self._first_index[v] = i

            if v == target:  # check only after all ops for this item
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

    # Filter only
    la6 = LazyArray([10, 20, 30, 40, 50])
    f6 = la6.filter(lambda n: n % 20 == 0)
    assert f6.indexOf(20) == 1
    assert f6.indexOf(40) == 3
    assert f6.indexOf(10) == -1

    # Map then filter
    la7 = LazyArray([1, 2, 3, 4, 5])
    m7 = la7.map(lambda n: n * 3).filter(lambda n: n % 2 == 0)
    assert m7.indexOf(6) == 1
    assert m7.indexOf(12) == 3
    assert m7.indexOf(9) == -1

    # Filter then map
    la8 = LazyArray([1, 2, 3, 4, 5])
    m8 = la8.filter(lambda n: n % 2 == 1).map(lambda n: n * 10)
    assert m8.indexOf(10) == 0
    assert m8.indexOf(50) == 4
    assert m8.indexOf(20) == -1

    # Early termination: should stop once target is found and not process later values.
    processed: list[int] = []

    def map_count(x: int) -> int:
        if x > 5:
            raise AssertionError("Should NOT reach value 6")
        processed.append(x)
        return x

    la9 = LazyArray([3, 4, 5, 6]).filter(lambda n: n % 2 == 1).map(map_count)
    assert la9.indexOf(5) == 2
    assert processed == [3, 5]

    # Repeated indexOf (original version) recomputes.
    seen_plain: list[int] = []

    def expensive_plain(x: int) -> int:
        seen_plain.append(x)
        return x

    la10 = LazyArray([1, 2, 3]).map(expensive_plain)
    assert la10.indexOf(2) == 1
    assert seen_plain == [1, 2]
    assert la10.indexOf(2) == 1
    assert seen_plain == [1, 2, 1, 2]

    # Repeated indexOf2 should reuse cached prefix results.
    seen: list[int] = []

    def expensive(x: int) -> int:
        seen.append(x)
        return x

    la11 = LazyArray([1, 2, 3]).map(expensive)
    assert la11.indexOf2(2) == 1
    assert seen == [1, 2]
    assert la11.indexOf2(3) == 2
    assert seen == [1, 2, 3]
    assert la11.indexOf2(3) == 2
    assert seen == [1, 2, 3]

    print("All tests passed!")
