from typing import List, Optional, Tuple


class QuerySystem:
    def __init__(self, records: List[List[str]]):
        parsed = []
        for row in records:
            if len(row) < 5:
                continue
            try:
                timestamp = int(row[0])
                amount = int(row[4])
            except ValueError:
                continue
            parsed.append(
                {
                    "raw": row[:],
                    "timestamp": timestamp,
                    "user_id": row[2],
                    "currency": row[3],
                    "amount": amount,
                }
            )
        self._records = sorted(parsed, key=lambda item: item["timestamp"])
        self._page_size = 0
        self._cursor = 0
        self._time_range: Optional[Tuple[int, int]] = None
        self._amount_range: Optional[Tuple[int, int]] = None
        self._user_id: Optional[str] = None
        self._currency: Optional[str] = None

    def setPageSize(self, size: int) -> None:
        self._page_size = size
        self._cursor = 0

    def setTimeRange(self, start: int, end: int) -> None:
        if start > end:
            start, end = end, start
        self._time_range = (start, end)
        self._cursor = 0

    def setAmountRange(self, start: int, end: int) -> None:
        if start > end:
            start, end = end, start
        self._amount_range = (start, end)
        self._cursor = 0

    def setUserId(self, user_id: str) -> None:
        self._user_id = user_id
        self._cursor = 0

    def setCurrency(self, currency: str) -> None:
        self._currency = currency
        self._cursor = 0

    def nextPage(self) -> List[List[str]]:
        limit = self._page_size if self._page_size > 0 else len(self._records)
        result: List[List[str]] = []
        while self._cursor < len(self._records) and len(result) < limit:
            current = self._records[self._cursor]
            self._cursor += 1
            if not self._matches(current):
                continue
            result.append(current["raw"][:])
        return result

    def _matches(self, record) -> bool:
        if self._time_range is not None:
            if not (self._time_range[0] <= record["timestamp"] <= self._time_range[1]):
                return False
        if self._amount_range is not None:
            if not (self._amount_range[0] <= record["amount"] <= self._amount_range[1]):
                return False
        if self._user_id is not None and record["user_id"] != self._user_id:
            return False
        if self._currency is not None and record["currency"] != self._currency:
            return False
        return True


class QueryPaginationSystem(QuerySystem):
    pass


if __name__ == "__main__":
    records = [
        ["1", "id-1", "user-1", "USD", "5"],
        ["2", "id-2", "user-2", "USD", "10"],
        ["3", "id-3", "user-1", "CAD", "20"],
        ["4", "id-4", "user-1", "CAD", "10"],
        ["5", "id-5", "user-1", "AUD", "30"],
        ["6", "id-6", "user-1", "JPY", "100"],
    ]

    system = QuerySystem(records)
    system.setPageSize(2)
    system.setTimeRange(1, 5)
    system.setUserId("user-1")
    assert system.nextPage() == [
        ["1", "id-1", "user-1", "USD", "5"],
        ["3", "id-3", "user-1", "CAD", "20"],
    ]
    assert system.nextPage() == [
        ["4", "id-4", "user-1", "CAD", "10"],
        ["5", "id-5", "user-1", "AUD", "30"],
    ]
    assert system.nextPage() == []

    unsorted = [
        ["10", "id-10", "u1", "USD", "1"],
        ["2", "id-2", "u1", "USD", "1"],
        ["7", "id-7", "u1", "USD", "1"],
    ]
    ordered = QueryPaginationSystem(unsorted)
    assert ordered.nextPage() == [
        ["2", "id-2", "u1", "USD", "1"],
        ["7", "id-7", "u1", "USD", "1"],
        ["10", "id-10", "u1", "USD", "1"],
    ]
