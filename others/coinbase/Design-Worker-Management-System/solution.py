from typing import Dict, List, Optional, Tuple


class WorkSession:
    def __init__(self, start: int, end: int, compensation: int) -> None:
        self.start = start
        self.end = end
        self.compensation = compensation


class _Worker:
    __slots__ = (
        "position",
        "compensation",
        "in_office",
        "last_entry",
        "total_time",
        "work_times",
        "sessions",
        "has_promotion",
        "promotion_start",
        "promotion_position",
        "promotion_compensation",
    )

    def __init__(self, position: str, compensation: int) -> None:
        self.position = position
        self.compensation = compensation
        self.in_office = False
        self.last_entry = 0
        self.total_time = 0
        self.work_times: Dict[str, int] = {}
        self.sessions: List[WorkSession] = []
        self.has_promotion = False
        self.promotion_start = 0
        self.promotion_position = ""
        self.promotion_compensation = 0


class OfficeManager:
    def __init__(self) -> None:
        self._workers: Dict[str, _Worker] = {}
        self._double_paid: List[Tuple[int, int]] = []

    def addWorker(self, workerId: str, position: str, compensation: int) -> bool:
        if workerId in self._workers:
            return False
        self._workers[workerId] = _Worker(position, compensation)
        return True

    def registerWorker(self, workerId: str, timestamp: int) -> str:
        worker = self._workers.get(workerId)
        if worker is None:
            return "invalid_request"

        if not worker.in_office:
            if worker.has_promotion and timestamp >= worker.promotion_start:
                worker.position = worker.promotion_position
                worker.compensation = worker.promotion_compensation
                worker.has_promotion = False

            worker.in_office = True
            worker.last_entry = timestamp
            return "registered"

        worker.in_office = False
        duration = timestamp - worker.last_entry
        worker.total_time += duration
        worker.work_times[worker.position] = (
            worker.work_times.get(worker.position, 0) + duration
        )
        worker.sessions.append(
            WorkSession(worker.last_entry, timestamp, worker.compensation)
        )
        return "registered"

    def get(self, workerId: str) -> int:
        worker = self._workers.get(workerId)
        if worker is None:
            return -1
        return worker.total_time

    def topNWorkers(self, n: int, position: str) -> str:
        if n <= 0:
            return ""
        entries = []
        for worker_id, worker in self._workers.items():
            if worker.position != position:
                continue
            entries.append((worker_id, worker.work_times.get(position, 0)))

        if not entries:
            return ""

        entries.sort(key=lambda x: (-x[1], x[0]))
        entries = entries[:n]
        return ", ".join(f"{worker_id}({time})" for worker_id, time in entries)

    def promote(
        self, workerId: str, newPosition: str, newCompensation: str, startTimestamp: int
    ) -> str:
        worker = self._workers.get(workerId)
        if worker is None:
            return "invalid_request"
        if worker.has_promotion:
            return "invalid_request"

        worker.has_promotion = True
        worker.promotion_start = startTimestamp
        worker.promotion_position = newPosition
        worker.promotion_compensation = int(newCompensation)
        return "success"

    def calcSalary(self, workerId: str, startTimestamp: int, endTimestamp: int) -> int:
        worker = self._workers.get(workerId)
        if worker is None:
            return -1

        total = 0
        for session in worker.sessions:
            session_start = max(session.start, startTimestamp)
            session_end = min(session.end, endTimestamp)
            if session_start >= session_end:
                continue
            duration = session_end - session_start
            total += duration * session.compensation
            total += (
                self._double_paid_overlap(session_start, session_end)
                * session.compensation
            )
        return total

    def setDoublePaid(self, startTimestamp: int, endTimestamp: int) -> None:
        if startTimestamp >= endTimestamp:
            return
        self._double_paid.append((startTimestamp, endTimestamp))
        self._double_paid.sort()

        merged: List[Tuple[int, int]] = []
        for start, end in self._double_paid:
            if not merged:
                merged.append((start, end))
                continue
            last_start, last_end = merged[-1]
            if start <= last_end:
                merged[-1] = (last_start, max(last_end, end))
            else:
                merged.append((start, end))
        self._double_paid = merged

    def _double_paid_overlap(self, start: int, end: int) -> int:
        if start >= end:
            return 0
        total = 0
        for period_start, period_end in self._double_paid:
            if period_end <= start:
                continue
            if period_start >= end:
                break
            overlap_start = max(start, period_start)
            overlap_end = min(end, period_end)
            if overlap_start < overlap_end:
                total += overlap_end - overlap_start
        return total


def _run(ops: List[str], args: List[List[object]]) -> List[object]:
    manager: Optional[OfficeManager] = None
    out: List[object] = []
    for op, arg in zip(ops, args):
        if op == "OfficeManager":
            manager = OfficeManager()
            out.append(None)
        elif op == "addWorker":
            out.append(manager.addWorker(arg[0], arg[1], arg[2]))
        elif op == "registerWorker":
            out.append(manager.registerWorker(arg[0], arg[1]))
        elif op == "get":
            out.append(manager.get(arg[0]))
        elif op == "topNWorkers":
            out.append(manager.topNWorkers(arg[0], arg[1]))
        elif op == "promote":
            out.append(manager.promote(arg[0], arg[1], arg[2], arg[3]))
        elif op == "calcSalary":
            out.append(manager.calcSalary(arg[0], arg[1], arg[2]))
        elif op == "setDoublePaid":
            manager.setDoublePaid(arg[0], arg[1])
            out.append(None)
        else:
            raise ValueError(f"unknown operation: {op}")
    return out


if __name__ == "__main__":
    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "get",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "get",
            "get",
            "registerWorker",
        ],
        [
            [],
            ["Ashley", "Middle Developer", 150],
            ["Ashley", "Junior Developer", 100],
            ["Ashley", 10],
            ["Ashley", 25],
            ["Ashley"],
            ["Ashley", 40],
            ["Ashley", 67],
            ["Ashley", 100],
            ["Ashley"],
            ["Walter"],
            ["Walter", 120],
        ],
    ) == [
        None,
        True,
        False,
        "registered",
        "registered",
        15,
        "registered",
        "registered",
        "registered",
        42,
        -1,
        "invalid_request",
    ]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "get",
            "registerWorker",
            "registerWorker",
            "get",
        ],
        [
            [],
            ["John", "Senior Developer", 200],
            ["Ashely"],
            ["John", 15],
            ["John", 30],
            ["John"],
        ],
    ) == [None, True, -1, "registered", "registered", 15]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "get",
            "addWorker",
        ],
        [
            [],
            ["Oliver", "Middle Developer", 150],
            ["Oliver", 25],
            ["Oliver", 55],
            ["Oliver"],
            ["Oliver", "Middle Developer", 150],
        ],
    ) == [None, True, "registered", "registered", 30, False]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "get",
        ],
        [
            [],
            ["Sophia", "Senior Developer", 200],
            ["Sophia", 30],
            ["Sophia", 60],
            ["Sophia", 90],
            ["Sophia"],
        ],
    ) == [None, True, "registered", "registered", "registered", 30]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "get",
            "registerWorker",
            "registerWorker",
            "get",
        ],
        [
            [],
            ["Emma", "Junior Developer", 100],
            ["Emma", 20],
            ["Emma", 50],
            ["Emma", 60],
            ["Emma"],
            ["Emma", 80],
            ["Emma", 100],
            ["Emma"],
        ],
    ) == [
        None,
        True,
        "registered",
        "registered",
        "registered",
        30,
        "registered",
        "registered",
        50,
    ]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "addWorker",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "topNWorkers",
            "topNWorkers",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "topNWorkers",
            "topNWorkers",
        ],
        [
            [],
            ["John", "Junior Developer", 120],
            ["Jason", "Junior Developer", 120],
            ["Ashley", "Junior Developer", 120],
            ["John", 100],
            ["John", 150],
            ["Jason", 200],
            ["Jason", 250],
            ["Jason", 275],
            [5, "Junior Developer"],
            [1, "Junior Developer"],
            ["Ashley", 400],
            ["Ashley", 500],
            ["Jason", 575],
            [5, "Junior Developer"],
            [5, "Middle Developer"],
        ],
    ) == [
        None,
        True,
        True,
        True,
        "registered",
        "registered",
        "registered",
        "registered",
        "registered",
        "Jason(50), John(50), Ashley(0)",
        "Jason(50)",
        "registered",
        "registered",
        "registered",
        "Jason(350), Ashley(100), John(50)",
        "",
    ]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "promote",
            "registerWorker",
            "promote",
            "registerWorker",
            "registerWorker",
            "calcSalary",
            "topNWorkers",
            "registerWorker",
            "get",
            "topNWorkers",
            "topNWorkers",
            "calcSalary",
            "calcSalary",
        ],
        [
            [],
            ["John", "Middle Developer", 200],
            ["John", 100],
            ["John", 125],
            ["John", "Senior Developer", "500", 200],
            ["John", 150],
            ["John", "Senior Developer", "350", 250],
            ["John", 300],
            ["John", 325],
            ["John", 0, 500],
            [3, "Senior Developer"],
            ["John", 400],
            ["John"],
            [10, "Senior Developer"],
            [10, "Middle Developer"],
            ["John", 110, 350],
            ["John", 900, 1400],
        ],
    ) == [
        None,
        True,
        "registered",
        "registered",
        "success",
        "registered",
        "invalid_request",
        "registered",
        "registered",
        35000,
        "John(0)",
        "registered",
        250,
        "John(75)",
        "",
        45500,
        0,
    ]

    assert _run(
        [
            "OfficeManager",
            "addWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "registerWorker",
            "setDoublePaid",
            "setDoublePaid",
            "setDoublePaid",
            "calcSalary",
            "calcSalary",
        ],
        [
            [],
            ["John", "Middle Developer", 100],
            ["John", 100],
            ["John", 200],
            ["John", 500],
            ["John", 600],
            ["John", 900],
            ["John", 1000],
            [50, 170],
            [530, 650],
            [580, 900],
            ["John", 0, 250],
            ["John", 0, 1500],
        ],
    ) == [
        None,
        True,
        "registered",
        "registered",
        "registered",
        "registered",
        "registered",
        "registered",
        None,
        None,
        None,
        17000,
        44000,
    ]
