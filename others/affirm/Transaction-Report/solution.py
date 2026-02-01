from dataclasses import dataclass
from enum import Enum
from typing import Dict, List
import datetime as dt
import math


class Status(Enum):
    COMPLETED = "COMPLETED"
    PENDING = "PENDING"
    FAILED = "FAILED"


@dataclass
class Transaction:
    amount: int
    merchant_id: str
    status: Status
    user_id: str
    start_time: dt.datetime


def hour_diff(now: dt.datetime, start_time: dt.datetime) -> int:
    # Helper function (provided by interviewer in real setting).
    return int(math.floor((now - start_time).total_seconds() / 3600.0))


def build_last_24h_report(
    transactions: List[Transaction], now: dt.datetime
) -> Dict[str, Dict[int, int]]:
    report: Dict[str, Dict[int, int]] = {}

    for tx in transactions:
        if tx.status != Status.COMPLETED:
            continue

        diff = hour_diff(now, tx.start_time)
        if diff < 0 or diff >= 24:
            continue

        inner = report.get(tx.merchant_id)
        if inner is None:
            inner = {}
            report[tx.merchant_id] = inner
        inner[diff] = inner.get(diff, 0) + tx.amount

    return report


if __name__ == "__main__":
    now = dt.datetime(2025, 12, 2, 0, 0, 0)
    source1 = [
        Transaction(20, "M1", Status.COMPLETED, "U1", dt.datetime(2025, 12, 1, 1, 0, 0)),
        Transaction(5, "M1", Status.COMPLETED, "U2", dt.datetime(2025, 12, 1, 2, 0, 0)),
        Transaction(7, "M2", Status.COMPLETED, "U3", dt.datetime(2025, 12, 1, 23, 0, 0)),
    ]
    source2 = [
        Transaction(9, "M1", Status.COMPLETED, "U4", dt.datetime(2025, 12, 1, 1, 30, 0)),
        Transaction(99, "M1", Status.PENDING, "U5", dt.datetime(2025, 12, 1, 3, 0, 0)),
        Transaction(10, "M2", Status.COMPLETED, "U6", dt.datetime(2025, 12, 1, 0, 0, 0)),
        Transaction(11, "M3", Status.COMPLETED, "U7", dt.datetime(2025, 12, 2, 1, 0, 0)),
    ]

    report = build_last_24h_report(source1 + source2, now)
    print(report)
    assert report["M1"][23] == 20
    assert report["M1"][22] == 14
    assert report["M2"][1] == 7
