from typing import Dict, List, Optional, Tuple, Any


class ValueWithTTL:
    def __init__(self, value: str, set_timestamp: int, ttl: Optional[int]) -> None:
        self.value = value
        self.set_timestamp = set_timestamp
        self.ttl = ttl

    def is_alive_at(self, timestamp: int) -> bool:
        if self.ttl is None:
            return True
        return self.set_timestamp <= timestamp < self.set_timestamp + self.ttl


class Backup:
    def __init__(self, timestamp: int, snapshot: Dict[str, Dict[str, ValueWithTTL]]):
        self.timestamp = timestamp
        self.snapshot = snapshot


class InMemoryDB:
    def __init__(self) -> None:
        self._records: Dict[str, Dict[str, ValueWithTTL]] = {}
        self._backups: List[Backup] = []

    def setData(self, key: str, field: str, value: str) -> None:
        record = self._records.get(key)
        if record is None:
            record = {}
            self._records[key] = record
        record[field] = ValueWithTTL(value, 0, None)

    def getData(self, key: str, field: str) -> str:
        record = self._records.get(key)
        if record is None:
            return ""
        entry = record.get(field)
        return entry.value if entry is not None else ""

    def deleteData(self, key: str, field: str) -> bool:
        record = self._records.get(key)
        if record is None or field not in record:
            return False
        del record[field]
        if not record:
            del self._records[key]
        return True

    def scanData(self, key: str) -> List[str]:
        record = self._records.get(key)
        if record is None:
            return []
        fields = sorted(record.keys())
        return [f"{field}({record[field].value})" for field in fields]

    def scanDataByPrefix(self, key: str, prefix: str) -> List[str]:
        record = self._records.get(key)
        if record is None:
            return []
        fields = sorted(field for field in record.keys() if field.startswith(prefix))
        return [f"{field}({record[field].value})" for field in fields]

    def setDataAt(self, key: str, field: str, value: str, timestamp: int) -> None:
        record = self._records.get(key)
        if record is None:
            record = {}
            self._records[key] = record
        record[field] = ValueWithTTL(value, timestamp, None)

    def setDataAtWithTtl(
        self, key: str, field: str, value: str, timestamp: int, ttl: int
    ) -> None:
        record = self._records.get(key)
        if record is None:
            record = {}
            self._records[key] = record
        record[field] = ValueWithTTL(value, timestamp, ttl)

    def deleteDataAt(self, key: str, field: str, timestamp: int) -> bool:
        record = self._records.get(key)
        if record is None:
            return False
        entry = record.get(field)
        if entry is None or not entry.is_alive_at(timestamp):
            return False
        del record[field]
        if not record:
            del self._records[key]
        return True

    def getDataAt(self, key: str, field: str, timestamp: int) -> str:
        record = self._records.get(key)
        if record is None:
            return ""
        entry = record.get(field)
        if entry is not None and entry.is_alive_at(timestamp):
            return entry.value
        return ""

    def scanDataAt(self, key: str, timestamp: int) -> List[str]:
        record = self._records.get(key)
        if record is None:
            return []
        fields = sorted(
            field for field, entry in record.items() if entry.is_alive_at(timestamp)
        )
        return [f"{field}({record[field].value})" for field in fields]

    def scanDataByPrefixAt(self, key: str, prefix: str, timestamp: int) -> List[str]:
        record = self._records.get(key)
        if record is None:
            return []
        fields = sorted(
            field
            for field, entry in record.items()
            if field.startswith(prefix) and entry.is_alive_at(timestamp)
        )
        return [f"{field}({record[field].value})" for field in fields]

    def backup(self, timestamp: int) -> int:
        snapshot: Dict[str, Dict[str, ValueWithTTL]] = {}
        count = 0
        for key, record in self._records.items():
            new_record: Dict[str, ValueWithTTL] = {}
            for field, entry in record.items():
                if not entry.is_alive_at(timestamp):
                    continue
                if entry.ttl is None:
                    new_record[field] = ValueWithTTL(entry.value, timestamp, None)
                else:
                    remaining = entry.set_timestamp + entry.ttl - timestamp
                    if remaining > 0:
                        new_record[field] = ValueWithTTL(
                            entry.value, timestamp, remaining
                        )
            if new_record:
                snapshot[key] = new_record
                count += 1
        self._backups.append(Backup(timestamp, snapshot))
        return count

    def restore(self, timestamp: int, timestamp_to_restore: int) -> None:
        chosen: Optional[Backup] = None
        for backup in reversed(self._backups):
            if backup.timestamp <= timestamp_to_restore:
                chosen = backup
                break
        if chosen is None:
            return
        self._records = {}
        for key, record in chosen.snapshot.items():
            new_record: Dict[str, ValueWithTTL] = {}
            for field, entry in record.items():
                new_record[field] = ValueWithTTL(entry.value, timestamp, entry.ttl)
            self._records[key] = new_record


def _run(ops: List[str], args: List[List[Any]]) -> List[object]:
    db: Optional[InMemoryDB] = None
    out: List[object] = []
    for op, arg in zip(ops, args):
        if op == "InMemoryDB":
            db = InMemoryDB()
            out.append(None)
        elif op == "setData":
            db.setData(arg[0], arg[1], arg[2])
            out.append(None)
        elif op == "getData":
            out.append(db.getData(arg[0], arg[1]))
        elif op == "deleteData":
            out.append(db.deleteData(arg[0], arg[1]))
        elif op == "scanData":
            out.append(db.scanData(arg[0]))
        elif op == "scanDataByPrefix":
            out.append(db.scanDataByPrefix(arg[0], arg[1]))
        elif op == "setDataAt":
            db.setDataAt(arg[0], arg[1], arg[2], arg[3])
            out.append(None)
        elif op == "setDataAtWithTtl":
            db.setDataAtWithTtl(arg[0], arg[1], arg[2], arg[3], arg[4])
            out.append(None)
        elif op == "deleteDataAt":
            out.append(db.deleteDataAt(arg[0], arg[1], arg[2]))
        elif op == "getDataAt":
            out.append(db.getDataAt(arg[0], arg[1], arg[2]))
        elif op == "scanDataAt":
            out.append(db.scanDataAt(arg[0], arg[1]))
        elif op == "scanDataByPrefixAt":
            out.append(db.scanDataByPrefixAt(arg[0], arg[1], arg[2]))
        elif op == "backup":
            out.append(db.backup(arg[0]))
        elif op == "restore":
            db.restore(arg[0], arg[1])
            out.append(None)
        else:
            raise ValueError(f"unknown operation: {op}")
    return out


if __name__ == "__main__":
    assert _run(
        ["InMemoryDB", "setData", "setData", "setData", "scanDataByPrefix", "scanData"],
        [[], ["A", "BC", "E"], ["A", "BD", "F"], ["A", "C", "G"], ["A", "B"], ["A"]],
    ) == [None, None, None, None, ["BC(E)", "BD(F)"], ["BC(E)", "BD(F)", "C(G)"]]

    assert _run(
        [
            "InMemoryDB",
            "setDataAtWithTtl",
            "setDataAtWithTtl",
            "setDataAt",
            "scanDataByPrefixAt",
            "scanDataByPrefixAt",
        ],
        [
            [],
            ["A", "BC", "E", 1, 9],
            ["A", "BC", "E", 5, 10],
            ["A", "BD", "F", 5],
            ["A", "", 14],
            ["A", "", 15],
        ],
    ) == [None, None, None, None, ["BC(E)", "BD(F)"], ["BD(F)"]]

    assert _run(
        [
            "InMemoryDB",
            "setDataAt",
            "setDataAtWithTtl",
            "getDataAt",
            "setDataAtWithTtl",
            "scanDataAt",
            "scanDataAt",
            "scanDataAt",
            "deleteDataAt",
        ],
        [
            [],
            ["A", "B", "C", 1],
            ["X", "Y", "Z", 2, 15],
            ["X", "Y", 3],
            ["A", "D", "E", 4, 10],
            ["A", 13],
            ["X", 16],
            ["X", 17],
            ["X", "Y", 20],
        ],
    ) == [None, None, None, "Z", None, ["B(C)", "D(E)"], ["Y(Z)"], [], False]

    assert _run(
        [
            "InMemoryDB",
            "setDataAtWithTtl",
            "backup",
            "setDataAt",
            "backup",
            "deleteDataAt",
            "backup",
            "restore",
            "backup",
            "scanDataAt",
            "scanDataAt",
        ],
        [
            [],
            ["A", "B", "C", 1, 10],
            [3],
            ["A", "D", "E", 4],
            [5],
            ["A", "B", 8],
            [9],
            [10, 7],
            [11],
            ["A", 15],
            ["A", 16],
        ],
    ) == [None, None, 1, None, 1, True, 1, None, 1, ["B(C)", "D(E)"], ["D(E)"]]
from typing import List, Optional


class InMemoryDB:
    def __init__(self) -> None:
        self._records = {}

    def setData(self, key: str, field: str, value: str) -> None:
        record = self._records.get(key)
        if record is None:
            record = {}
            self._records[key] = record
        record[field] = value

    def getData(self, key: str, field: str) -> str:
        record = self._records.get(key)
        if record is None:
            return ""
        return record.get(field, "")

    def deleteData(self, key: str, field: str) -> bool:
        record = self._records.get(key)
        if record is None or field not in record:
            return False
        del record[field]
        if not record:
            del self._records[key]
        return True


def _run(ops: List[str], args: List[List[str]]) -> List[object]:
    db: Optional[InMemoryDB] = None
    out: List[object] = []
    for op, arg in zip(ops, args):
        if op == "InMemoryDB":
            db = InMemoryDB()
            out.append(None)
        elif op == "setData":
            db.setData(arg[0], arg[1], arg[2])
            out.append(None)
        elif op == "getData":
            out.append(db.getData(arg[0], arg[1]))
        elif op == "deleteData":
            out.append(db.deleteData(arg[0], arg[1]))
        else:
            raise ValueError(f"unknown operation: {op}")
    return out


if __name__ == "__main__":
    assert _run(
        [
            "InMemoryDB",
            "setData",
            "setData",
            "getData",
            "getData",
            "deleteData",
            "deleteData",
        ],
        [
            [],
            ["A", "B", "E"],
            ["A", "C", "F"],
            ["A", "B"],
            ["A", "D"],
            ["A", "B"],
            ["A", "D"],
        ],
    ) == [None, None, None, "E", "", True, False]

    assert _run(
        [
            "InMemoryDB",
            "setData",
            "setData",
            "getData",
            "getData",
            "deleteData",
            "getData",
            "getData",
        ],
        [
            [],
            ["user1", "name", "John"],
            ["user1", "age", "25"],
            ["user1", "name"],
            ["user1", "age"],
            ["user1", "age"],
            ["user1", "age"],
            ["user1", "name"],
        ],
    ) == [None, None, None, "John", "25", True, "", "John"]

    assert _run(
        [
            "InMemoryDB",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "getData",
            "getData",
            "getData",
            "setData",
            "getData",
            "deleteData",
            "getData",
            "getData",
        ],
        [
            [],
            ["record1", "field1", "value1_1"],
            ["record1", "field2", "value2_1"],
            ["record1", "field3", "value3_1"],
            ["record50", "field1", "value1_50"],
            ["record50", "field2", "value2_50"],
            ["record50", "field3", "value3_50"],
            ["record100", "field1", "value1_100"],
            ["record100", "field2", "value2_100"],
            ["record100", "field3", "value3_100"],
            ["record50", "field2"],
            ["record1", "field1"],
            ["record100", "field3"],
            ["record50", "field2", "updated_value"],
            ["record50", "field2"],
            ["record1", "field1"],
            ["record1", "field1"],
            ["record1", "field2"],
        ],
    ) == [
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        "value2_50",
        "value1_1",
        "value3_100",
        None,
        "updated_value",
        True,
        "",
        "value2_1",
    ]

    assert _run(
        [
            "InMemoryDB",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "setData",
            "getData",
            "getData",
            "getData",
            "getData",
            "deleteData",
            "deleteData",
            "getData",
            "getData",
        ],
        [
            [],
            ["largeRecord", "field9", "initialValue9"],
            ["largeRecord", "field10", "initialValue10"],
            ["largeRecord", "field100", "initialValue100"],
            ["largeRecord", "field250", "initialValue250"],
            ["largeRecord", "field750", "initialValue750"],
            ["largeRecord", "field999", "initialValue999"],
            ["largeRecord", "field10", "overwrittenValue10"],
            ["largeRecord", "field100", "overwrittenValue100"],
            ["largeRecord", "field10"],
            ["largeRecord", "field9"],
            ["largeRecord", "field100"],
            ["largeRecord", "field999"],
            ["largeRecord", "field250"],
            ["largeRecord", "field750"],
            ["largeRecord", "field250"],
            ["largeRecord", "field750"],
        ],
    ) == [
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        "overwrittenValue10",
        "initialValue9",
        "overwrittenValue100",
        "initialValue999",
        True,
        True,
        "",
        "",
    ]

    assert _run(
        [
            "InMemoryDB",
            "getData",
            "deleteData",
            "setData",
            "getData",
            "deleteData",
            "setData",
            "getData",
            "setData",
            "deleteData",
            "getData",
            "deleteData",
            "getData",
            "setData",
            "getData",
        ],
        [
            [],
            ["nonExistent", "field"],
            ["nonExistent", "field"],
            ["", "", ""],
            ["", ""],
            ["", ""],
            ["testKey", "nullField", "null"],
            ["testKey", "nullField"],
            ["existingKey", "field1", "value1"],
            ["existingKey", "nonExistentField"],
            ["existingKey", "field1"],
            ["existingKey", "field1"],
            ["existingKey", "field1"],
            ["key@#$%", "field!@#", "value~`"],
            ["key@#$%", "field!@#"],
        ],
    ) == [
        None,
        "",
        False,
        None,
        "",
        True,
        None,
        "null",
        None,
        False,
        "value1",
        True,
        "",
        None,
        "value~`",
    ]
