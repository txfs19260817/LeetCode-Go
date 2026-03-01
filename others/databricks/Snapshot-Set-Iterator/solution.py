from __future__ import annotations


class SnapshotSet:
    """A set of integers that supports snapshotted iteration in insertion order."""

    def __init__(self) -> None:
        # Single dict storage:
        # ("ent", add_version) -> (value, remove_version)
        # ("idx", value) -> latest add_version for this value
        self.store: dict[tuple[str, int], int | tuple[int, int]] = {}
        self.version: int = 0

    def add(self, n: int) -> bool:
        idx_key = ("idx", n)
        if idx_key in self.store:
            add_v = self.store[idx_key]
            _, rem_v = self.store[("ent", add_v)]
            if rem_v == -1:
                return False  # already present

        add_v = self.version
        self.store[("ent", add_v)] = (n, -1)
        self.store[idx_key] = add_v
        self.version += 1
        return True

    def remove(self, n: int) -> bool:
        idx_key = ("idx", n)
        if idx_key not in self.store:
            return False

        add_v = self.store[idx_key]
        val, rem_v = self.store[("ent", add_v)]
        if rem_v != -1:
            return False  # already removed

        self.store[("ent", add_v)] = (val, self.version)
        self.version += 1
        return True

    def contains(self, n: int) -> bool:
        idx_key = ("idx", n)
        if idx_key not in self.store:
            return False

        add_v = self.store[idx_key]
        _, rem_v = self.store[("ent", add_v)]
        return rem_v == -1

    def get_iterator(self) -> SnapshotIterator:
        return SnapshotIterator(
            snapshot_set=self,
            snap_version=self.version,
            scan_limit=self.version,
        )


class SnapshotIterator:
    """Iterates over a frozen snapshot of the set."""

    def __init__(
        self, snapshot_set: SnapshotSet, snap_version: int, scan_limit: int
    ) -> None:
        self._set = snapshot_set
        self._snap_version = snap_version
        self._scan_limit = scan_limit
        self._scan_pos = 0

    def next(self) -> int:
        while self._scan_pos < self._scan_limit:
            add_v = self._scan_pos
            self._scan_pos += 1
            entry = self._set.store.get(("ent", add_v))
            if entry is None:
                continue
            val, rem_v = entry
            if rem_v == -1 or rem_v >= self._snap_version:
                return val
        raise StopIteration("no next element")


if __name__ == "__main__":
    def drain(it: SnapshotIterator) -> list[int]:
        res = []
        while True:
            try:
                res.append(it.next())
            except StopIteration:
                break
        return res

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

    result1 = drain(it1)
    assert result1 == [1, 2, 3, 4], f"Expected [1,2,3,4], got {result1}"

    result2 = drain(it2)
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
    result = drain(it)
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

    assert drain(it_a) == [10, 20]
    assert drain(it_b) == [10, 20, 30]
    assert drain(it_c) == [20, 30]

    # Test 4: Iterator on empty set
    s4 = SnapshotSet()
    it_empty = s4.get_iterator()
    assert drain(it_empty) == []

    print("All tests passed!")
