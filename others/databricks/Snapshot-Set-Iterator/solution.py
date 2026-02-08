from __future__ import annotations


class SnapshotSet:
    """A set of integers that supports snapshotted iteration in insertion order."""

    def __init__(self) -> None:
        # Each entry: (value, add_version, remove_version)  remove_version == -1 means active
        self.entries: list[tuple[int, int, int]] = []
        self.index: dict[int, int] = {}  # value -> index in entries
        self.version: int = 0

    def add(self, n: int) -> bool:
        if n in self.index:
            idx = self.index[n]
            _, _, rem_v = self.entries[idx]
            if rem_v == -1:
                return False  # already present
        # Append new entry (handles both fresh add and re-add after remove)
        self.entries.append((n, self.version, -1))
        self.index[n] = len(self.entries) - 1
        self.version += 1
        return True

    def remove(self, n: int) -> bool:
        if n not in self.index:
            return False
        idx = self.index[n]
        val, add_v, rem_v = self.entries[idx]
        if rem_v != -1:
            return False  # already removed
        self.entries[idx] = (val, add_v, self.version)
        self.version += 1
        return True

    def contains(self, n: int) -> bool:
        if n not in self.index:
            return False
        _, _, rem_v = self.entries[self.index[n]]
        return rem_v == -1

    def get_iterator(self) -> SnapshotIterator:
        snap_version = self.version
        snapshot = [
            e[0]
            for e in self.entries
            if e[1] < snap_version and (e[2] == -1 or e[2] >= snap_version)
        ]
        return SnapshotIterator(snapshot)


class SnapshotIterator:
    """Iterates over a frozen snapshot of the set."""

    def __init__(self, elements: list[int]) -> None:
        self._elements = elements
        self._pos = 0

    def has_next(self) -> bool:
        return self._pos < len(self._elements)

    def next(self) -> int:
        if not self.has_next():
            raise StopIteration("no next element")
        val = self._elements[self._pos]
        self._pos += 1
        return val


if __name__ == "__main__":
    # Test 1: Main example
    s = SnapshotSet()
    assert s.add(1) is True
    assert s.add(2) is True
    assert s.add(3) is True
    assert s.add(4) is True
    assert s.add(1) is False  # already exists

    it1 = s.get_iterator()  # snapshot: [1,2,3,4]

    assert s.remove(1) is True
    assert s.remove(3) is True
    assert s.remove(5) is False

    it2 = s.get_iterator()  # snapshot: [2,4]

    result1 = []
    while it1.has_next():
        result1.append(it1.next())
    assert result1 == [1, 2, 3, 4], f"Expected [1,2,3,4], got {result1}"

    result2 = []
    while it2.has_next():
        result2.append(it2.next())
    assert result2 == [2, 4], f"Expected [2,4], got {result2}"

    # Test 2: Re-add after remove (appears at new position)
    s2 = SnapshotSet()
    assert s2.add(1) is True
    assert s2.add(2) is True
    assert s2.remove(1) is True
    assert s2.contains(1) is False
    assert s2.add(1) is True  # re-add
    assert s2.contains(1) is True

    it = s2.get_iterator()
    result = []
    while it.has_next():
        result.append(it.next())
    assert result == [2, 1], f"Expected [2,1], got {result}"

    # Test 3: Multiple concurrent iterators
    s3 = SnapshotSet()
    assert s3.add(10) is True
    assert s3.add(20) is True
    it_a = s3.get_iterator()  # [10, 20]

    assert s3.add(30) is True
    it_b = s3.get_iterator()  # [10, 20, 30]

    assert s3.remove(10) is True
    it_c = s3.get_iterator()  # [20, 30]

    def drain(it: SnapshotIterator) -> list[int]:
        res = []
        while it.has_next():
            res.append(it.next())
        return res

    assert drain(it_a) == [10, 20]
    assert drain(it_b) == [10, 20, 30]
    assert drain(it_c) == [20, 30]

    # Test 4: Iterator on empty set
    s4 = SnapshotSet()
    it_empty = s4.get_iterator()
    assert it_empty.has_next() is False
    assert drain(it_empty) == []

    print("All tests passed!")
