import heapq
from dataclasses import dataclass, field


@dataclass
class Customer:
    id: int
    total_revenue: int = field(default=0)


class RevenueSystem:
    def __init__(self) -> None:
        self.customers: list[Customer] = []

    def add(self, revenue: int) -> int:
        cid = len(self.customers)
        self.customers.append(Customer(id=cid, total_revenue=revenue))
        return cid

    def add_by_referral(self, revenue: int, referrer_id: int) -> int:
        if referrer_id < 0 or referrer_id >= len(self.customers):
            return -1
        cid = self.add(revenue)
        self.customers[referrer_id].total_revenue += revenue
        return cid

    def get_top_k_customer(self, k: int, min_revenue: int) -> list[int]:
        # Min-heap of (total_revenue, id); keep at most k entries.
        h: list[tuple[int, int]] = []
        for c in self.customers:
            if c.total_revenue >= min_revenue:
                if len(h) < k:
                    heapq.heappush(h, (c.total_revenue, c.id))
                elif c.total_revenue > h[0][0]:
                    heapq.heapreplace(h, (c.total_revenue, c.id))
        # Pop all and reverse to get descending order.
        result = [heapq.heappop(h)[1] for _ in range(len(h))]
        result.reverse()
        return result


if __name__ == "__main__":
    # Test 1: Main example
    rs = RevenueSystem()
    assert rs.add(100) == 0           # Customer 0: $100
    assert rs.add(50) == 1            # Customer 1: $50
    assert rs.add_by_referral(30, 0) == 2  # Customer 0: $130, Customer 2: $30
    assert rs.add_by_referral(70, 1) == 3  # Customer 1: $120, Customer 3: $70
    assert rs.add(50) == 4            # Customer 4: $50
    assert rs.get_top_k_customer(2, 100) == [0, 1]
    assert rs.add_by_referral(50, 4) == 5  # Customer 4: $100, Customer 5: $50
    assert rs.get_top_k_customer(3, 100) == [0, 1, 4]

    # Test 2: Invalid referrer
    rs2 = RevenueSystem()
    assert rs2.add(100) == 0
    assert rs2.add_by_referral(50, 5) == -1
    assert rs2.add_by_referral(50, -1) == -1
    assert rs2.get_top_k_customer(10, 0) == [0]

    # Test 3: Multiple referrals to same customer
    rs3 = RevenueSystem()
    assert rs3.add(100) == 0             # Customer 0: $100
    assert rs3.add_by_referral(50, 0) == 1  # Customer 0: $150, Customer 1: $50
    assert rs3.add_by_referral(30, 0) == 2  # Customer 0: $180, Customer 2: $30
    assert rs3.add_by_referral(20, 0) == 3  # Customer 0: $200, Customer 3: $20
    assert rs3.get_top_k_customer(5, 100) == [0]
    assert rs3.get_top_k_customer(10, 0) == [0, 1, 2, 3]
    assert rs3.get_top_k_customer(2, 0) == [0, 1]

    print("Base tests passed!")


# ===================================================================
# Follow-up: Nested Revenue (entire referral subtree)
# ===================================================================


class NestedRevenueReadHeavy:
    """Eager propagation: O(D) write, O(N log K) read."""

    def __init__(self) -> None:
        self._own: list[int] = []
        self._nested: list[int] = []
        self._parent: list[int] = []
        self._children: list[list[int]] = []

    def add(self, revenue: int) -> int:
        cid = len(self._own)
        self._own.append(revenue)
        self._nested.append(revenue)
        self._parent.append(-1)
        self._children.append([])
        return cid

    def add_by_referral(self, revenue: int, referrer_id: int) -> int:
        if referrer_id < 0 or referrer_id >= len(self._own):
            return -1
        cid = self.add(revenue)
        self._parent[cid] = referrer_id
        self._children[referrer_id].append(cid)
        # Propagate up — O(D)
        cur = referrer_id
        while cur != -1:
            self._nested[cur] += revenue
            cur = self._parent[cur]
        return cid

    def get_top_k_customer(self, k: int, min_revenue: int) -> list[int]:
        return _top_k(self._nested, k, min_revenue)


class NestedRevenueWriteHeavy:
    """Lazy aggregation: O(1) write, O(N log K) read."""

    def __init__(self) -> None:
        self._own: list[int] = []
        self._parent: list[int] = []
        self._children: list[list[int]] = []

    def add(self, revenue: int) -> int:
        cid = len(self._own)
        self._own.append(revenue)
        self._parent.append(-1)
        self._children.append([])
        return cid

    def add_by_referral(self, revenue: int, referrer_id: int) -> int:
        if referrer_id < 0 or referrer_id >= len(self._own):
            return -1
        cid = self.add(revenue)
        self._parent[cid] = referrer_id
        self._children[referrer_id].append(cid)
        return cid  # O(1)

    def get_top_k_customer(self, k: int, min_revenue: int) -> list[int]:
        n = len(self._own)
        nested = [0] * n
        for i in range(n - 1, -1, -1):
            nested[i] = self._own[i] + sum(nested[c] for c in self._children[i])
        return _top_k(nested, k, min_revenue)


def _top_k(revenues: list[int], k: int, min_revenue: int) -> list[int]:
    h: list[tuple[int, int]] = []
    for i, rev in enumerate(revenues):
        if rev < min_revenue:
            continue
        if len(h) < k:
            heapq.heappush(h, (rev, i))
        elif rev > h[0][0]:
            heapq.heapreplace(h, (rev, i))
    result = [heapq.heappop(h)[1] for _ in range(len(h))]
    result.reverse()
    return result


if __name__ == "__main__":
    # Nested: 0(100)→2(40)→3(25), 1(60)→4(10)
    # Nested revenues: 0=165, 1=70, 2=65, 3=25, 4=10
    for Cls in (NestedRevenueReadHeavy, NestedRevenueWriteHeavy):
        s = Cls()
        assert s.add(100) == 0
        assert s.add(60) == 1
        assert s.add_by_referral(40, 0) == 2
        assert s.add_by_referral(25, 2) == 3
        assert s.add_by_referral(10, 1) == 4
        assert s.get_top_k_customer(3, 0) == [0, 1, 2], f"{Cls.__name__}"
        assert s.get_top_k_customer(2, 50) == [0, 1], f"{Cls.__name__}"
        assert s.get_top_k_customer(5, 100) == [0], f"{Cls.__name__}"
        # Deep chain: 0=100, 1=90, 2=70, 3=40
        d = Cls()
        d.add(10)
        d.add_by_referral(20, 0)
        d.add_by_referral(30, 1)
        d.add_by_referral(40, 2)
        assert d.get_top_k_customer(10, 0) == [0, 1, 2, 3], f"{Cls.__name__}"
        # Invalid referrer
        e = Cls()
        e.add(100)
        assert e.add_by_referral(50, 5) == -1
        assert e.get_top_k_customer(10, 0) == [0]

    print("All tests passed!")
