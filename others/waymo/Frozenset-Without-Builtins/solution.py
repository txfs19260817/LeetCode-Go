def _lower_bound(values: tuple[int, ...], target: int) -> tuple[int, bool]:
    left, right = 0, len(values)
    while left < right:
        mid = left + (right - left) // 2
        if values[mid] < target:
            left = mid + 1
        else:
            right = mid
    if left < len(values) and values[left] == target:
        return left, True
    return left, False


class FrozenSet:
    __slots__ = ("_elems",)

    def __init__(self, elems: tuple[int, ...] = ()) -> None:
        self._elems = elems

    @classmethod
    def from_values(cls, values: list[int]) -> "FrozenSet":
        result = cls()
        for value in values:
            result = result.with_value(value)
        return result

    def size(self) -> int:
        return len(self._elems)

    def contains(self, value: int) -> bool:
        _, found = _lower_bound(self._elems, value)
        return found

    def elements(self) -> list[int]:
        return list(self._elems)

    def with_value(self, value: int) -> "FrozenSet":
        index, found = _lower_bound(self._elems, value)
        if found:
            return self
        next_elems = self._elems[:index] + (value,) + self._elems[index:]
        return FrozenSet(next_elems)

    def without_value(self, value: int) -> "FrozenSet":
        index, found = _lower_bound(self._elems, value)
        if not found:
            return self
        next_elems = self._elems[:index] + self._elems[index + 1 :]
        return FrozenSet(next_elems)

    def union(self, other: "FrozenSet") -> "FrozenSet":
        left, right = self._elems, other._elems
        out: list[int] = []
        i, j = 0, 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                out.append(left[i])
                i += 1
            elif left[i] > right[j]:
                out.append(right[j])
                j += 1
            else:
                out.append(left[i])
                i += 1
                j += 1
        while i < len(left):
            out.append(left[i])
            i += 1
        while j < len(right):
            out.append(right[j])
            j += 1
        return FrozenSet(tuple(out))

    def intersection(self, other: "FrozenSet") -> "FrozenSet":
        left, right = self._elems, other._elems
        out: list[int] = []
        i, j = 0, 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                i += 1
            elif left[i] > right[j]:
                j += 1
            else:
                out.append(left[i])
                i += 1
                j += 1
        return FrozenSet(tuple(out))

    def difference(self, other: "FrozenSet") -> "FrozenSet":
        left, right = self._elems, other._elems
        out: list[int] = []
        i, j = 0, 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                out.append(left[i])
                i += 1
            elif left[i] > right[j]:
                j += 1
            else:
                i += 1
                j += 1
        while i < len(left):
            out.append(left[i])
            i += 1
        return FrozenSet(tuple(out))

    def equals(self, other: "FrozenSet") -> bool:
        return self._elems == other._elems


if __name__ == "__main__":
    s1 = FrozenSet.from_values([5, 1, 3, 3])
    s2 = s1.with_value(4)
    s3 = s2.without_value(1)
    s4 = FrozenSet.from_values([3, 7])

    assert s1.elements() == [1, 3, 5]
    assert s2.elements() == [1, 3, 4, 5]
    assert s3.elements() == [3, 4, 5]
    assert s3.union(s4).elements() == [3, 4, 5, 7]
    assert s3.intersection(s4).elements() == [3]
    assert s3.difference(s4).elements() == [4, 5]

    assert s1.contains(1) is True
    assert s1.contains(2) is False
    assert s1.size() == 3
    assert s1.equals(FrozenSet.from_values([3, 5, 1])) is True

    chain = FrozenSet()
    for i in range(1, 201):
        chain = chain.with_value(i)
    assert chain.size() == 200
    assert chain.elements()[:5] == [1, 2, 3, 4, 5]
    assert chain.elements()[-5:] == [196, 197, 198, 199, 200]
