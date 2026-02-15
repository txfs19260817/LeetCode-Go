from __future__ import annotations

import math


class KVStoreWithQPS:
    """Key-value store that tracks average QPS over a 5-minute window.

    Uses a circular buffer of size 300 (one slot per second) so that
    space is O(W) regardless of throughput — the LC 362 optimisation.
    """

    WINDOW = 300

    def __init__(self) -> None:
        self.data: dict[str, str] = {}
        self.times: list[int] = [0] * self.WINDOW
        self.counts: list[int] = [0] * self.WINDOW

    def _hit(self, timestamp: int) -> None:
        idx = timestamp % self.WINDOW
        if self.times[idx] == timestamp:
            self.counts[idx] += 1
        else:
            self.times[idx] = timestamp
            self.counts[idx] = 1

    def put(self, key: str, value: str, timestamp: int) -> None:
        self.data[key] = value
        self._hit(timestamp)

    def get(self, key: str, timestamp: int) -> str:
        self._hit(timestamp)
        return self.data.get(key, "")

    def get_qps(self, timestamp: int) -> float:
        """Average queries-per-second over (timestamp-300, timestamp]."""
        total = 0
        for i in range(self.WINDOW):
            if timestamp - self.times[i] < self.WINDOW:
                total += self.counts[i]
        return total / self.WINDOW


if __name__ == "__main__":
    # Sample walkthrough: 3 hits / 300 = 0.01
    s = KVStoreWithQPS()
    s.put("foo", "bar", 1)
    s.put("baz", "qux", 2)
    assert s.get("foo", 3) == "bar"
    assert math.isclose(s.get_qps(3), 0.01)

    # Expiry: t=1 falls out at t=301 → 2/300
    s2 = KVStoreWithQPS()
    s2.put("a", "1", 1)
    s2.put("b", "2", 2)
    s2.put("c", "3", 301)
    assert math.isclose(s2.get_qps(301), 2 / 300)

    # All expired
    s3 = KVStoreWithQPS()
    s3.put("a", "1", 1)
    s3.put("b", "2", 2)
    assert math.isclose(s3.get_qps(500), 0.0, abs_tol=1e-9)

    # Multiple ops at same timestamp: 3/300 = 0.01
    s4 = KVStoreWithQPS()
    s4.put("a", "1", 5)
    s4.put("b", "2", 5)
    s4.put("c", "3", 5)
    assert math.isclose(s4.get_qps(5), 0.01)

    # Get on missing key still counts as hit: 1/300
    s5 = KVStoreWithQPS()
    assert s5.get("x", 1) == ""
    assert math.isclose(s5.get_qps(1), 1 / 300)

    # Overwrite: 2 puts + 1 get = 3/300 = 0.01
    s6 = KVStoreWithQPS()
    s6.put("k", "v1", 1)
    s6.put("k", "v2", 2)
    assert s6.get("k", 3) == "v2"
    assert math.isclose(s6.get_qps(3), 0.01)

    # Boundary: t=1 inside at t=300 (1/300), outside at t=301 (0)
    s7 = KVStoreWithQPS()
    s7.put("a", "1", 1)
    assert math.isclose(s7.get_qps(300), 1 / 300)
    assert math.isclose(s7.get_qps(301), 0.0, abs_tol=1e-9)

    # High throughput: 600 ops in 1 second → 600/300 = 2.0 QPS
    s8 = KVStoreWithQPS()
    for _ in range(600):
        s8.put("k", "v", 10)
    assert math.isclose(s8.get_qps(10), 2.0)

    print("All tests passed!")
