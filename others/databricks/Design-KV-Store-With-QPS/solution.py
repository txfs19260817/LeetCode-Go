from __future__ import annotations

import math
import threading
import time
from collections.abc import Callable


class KVStoreWithQPS:
    """Key-value store that tracks average QPS over a 5-minute window.

    Uses a circular buffer of size 300 (one slot per second) so that
    space is O(W) regardless of throughput — the LC 362 optimisation.
    """

    WINDOW = 300

    def __init__(self) -> None:
        """Initialize an empty store and fixed-size 300-second ring buffers."""
        self.data: dict[str, str] = {}
        # Slot i corresponds to second timestamps where ts % WINDOW == i.
        self.times: list[int] = [0] * self.WINDOW
        self.counts: list[int] = [0] * self.WINDOW

    def _hit(self, timestamp: int) -> None:
        """Record one operation at `timestamp` into the ring buffer."""
        idx = timestamp % self.WINDOW
        if self.times[idx] == timestamp:
            # Same second, aggregate count in-place.
            self.counts[idx] += 1
        else:
            # Slot reused by a newer second; old value is implicitly expired.
            self.times[idx] = timestamp
            self.counts[idx] = 1

    def put(self, key: str, value: str, timestamp: int) -> None:
        """Store `key -> value` and count this operation as one hit."""
        self.data[key] = value
        self._hit(timestamp)

    def get(self, key: str, timestamp: int) -> str:
        """Return value for `key` (or empty string), and count one hit."""
        self._hit(timestamp)
        return self.data.get(key, "")

    def get_qps(self, timestamp: int) -> float:
        """Average queries-per-second over (timestamp-300, timestamp]."""
        total = 0
        for i in range(self.WINDOW):
            # Keep only hits in the last WINDOW seconds.
            if timestamp - self.times[i] < self.WINDOW:
                total += self.counts[i]
        return total / self.WINDOW


class KVStoreRealtime:
    """Production-style KV store with real-time QPS tracking.

    Differences from KVStoreWithQPS:
    - Uses real clock time instead of injected timestamps.
    - Maintains a running total so get_qps is O(1).
    - Uses min(elapsed, window_seconds) as denominator during startup.
    - Thread-safe via a lock.
    """

    def __init__(self, window_seconds: int = 300, now_fn: Callable[[], int] | None = None) -> None:
        """Create a real-time store with configurable window and clock source."""
        if window_seconds <= 0:
            raise ValueError("window_seconds must be > 0")

        # now_fn is injectable to make deterministic tests easy.
        self._now_fn = now_fn if now_fn is not None else lambda: int(time.time())
        now_sec = self._now()

        self._lock = threading.Lock()
        self._data: dict[str, str] = {}
        self._window_sec = window_seconds
        self._buckets: list[int] = [0] * window_seconds
        self._total = 0
        self._last_sec = now_sec
        self._start_sec = now_sec

    def _now(self) -> int:
        """Return current unix second from the configured clock."""
        return int(self._now_fn())

    def _normalize(self, now_sec: int) -> None:
        """Advance time to `now_sec` and evict buckets outside the window."""
        if now_sec <= self._last_sec:
            return

        diff = now_sec - self._last_sec
        if diff >= self._window_sec:
            # Time jumped beyond one full window, so all buckets are stale.
            self._buckets = [0] * self._window_sec
            self._total = 0
            self._last_sec = now_sec
            return

        # Only clear slots for newly passed seconds.
        for sec in range(self._last_sec + 1, now_sec + 1):
            idx = sec % self._window_sec
            self._total -= self._buckets[idx]
            self._buckets[idx] = 0
        self._last_sec = now_sec

    def _record_op(self, now_sec: int) -> None:
        """Normalize first, then record one operation at `now_sec`."""
        self._normalize(now_sec)
        idx = now_sec % self._window_sec
        # Add one operation into the current second bucket.
        self._buckets[idx] += 1
        self._total += 1

    def put(self, key: str, value: str) -> None:
        """Store `key -> value` and count this call as an operation."""
        with self._lock:
            self._record_op(self._now())
            self._data[key] = value

    def get(self, key: str) -> str:
        """Return value for `key` (or empty string), and count one operation."""
        with self._lock:
            self._record_op(self._now())
            return self._data.get(key, "")

    def get_qps(self) -> float:
        """Return average QPS over the effective real-time sliding window."""
        # get_qps itself is not counted as an operation.
        with self._lock:
            now_sec = self._now()
            self._normalize(now_sec)

            # During startup, divide by real elapsed seconds instead of full window.
            elapsed = now_sec - self._start_sec + 1
            denom = min(elapsed, self._window_sec)
            if denom <= 0:
                return 0.0
            return self._total / denom


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

    # Realtime version (with controllable clock for deterministic testing).
    fake_now = [1000]

    def now_fn() -> int:
        return fake_now[0]

    r1 = KVStoreRealtime(window_seconds=5, now_fn=now_fn)
    r1.put("foo", "bar")
    assert r1.get("foo") == "bar"
    # Two operations in the first second -> denominator is 1 during startup.
    assert math.isclose(r1.get_qps(), 2.0)

    # Startup denominator ramps with elapsed time.
    fake_now[0] = 1002
    assert math.isclose(r1.get_qps(), 2 / 3)

    # After window passes, all old hits expire.
    fake_now[0] = 1006
    assert math.isclose(r1.get_qps(), 0.0, abs_tol=1e-9)

    print("All tests passed!")
