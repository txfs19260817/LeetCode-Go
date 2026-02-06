# Design In-memory Database with Backup

Implement a simplified version of an in-memory database to store records. Each record can be
accessed with a unique identifier key of string type. A record may contain several field-value
pairs, both of which are of string type.

Implement the InMemoryDB class:

`void setData(String key, String field, String value)` should insert a field-value pair to the
record associated with key.

- If the field in the record already exists, replace the existing value with the specified value.
- If the record does not exist, create a new one.

`String getData(String key, String field)` should return the value contained within the field of
the record associated with key.

- If the record or the field doesn't exist, it should return "".

`boolean deleteData(String key, String field)` should remove the field from the record associated
with key.

- Returns true if the field was successfully deleted, and false if the key or the field does not
  exist in the database.

## Constraints

- Total number of calls to all methods ≤ 10^5.
- 1 ≤ key.length, field.length, value.length ≤ 100.
- All strings contain printable ASCII characters.

## Example

**Input:**

```
["InMemoryDB", "setData", "setData", "getData", "getData", "deleteData", "deleteData"]
[[], ["A", "B", "E"], ["A", "C", "F"], ["A", "B"], ["A", "D"], ["A", "B"], ["A", "D"]]
```

**Output:**

```
[null, null, null, "E", "", true, false]
```

InMemoryDB db = new InMemoryDB();
db.setData("A", "B", "E"); // Returns null. Database state: {"A": {"B": "E"}}
db.setData("A", "C", "F"); // Returns null. Database state: {"A": {"C": "F", "B": "E"}}
db.getData("A", "B"); // Returns "E".
db.getData("A", "D"); // Returns "", since there is no value of field "D".
db.deleteData("A", "B"); // Returns true. Database state: {"A": {"C": "F"}}
db.deleteData("A", "D"); // Returns false. Database state: {"A": {"C": "F"}}

## Follow-up 1: Scan Operations

List<String> scanData(String key) should return a list of strings representing
the fields of the record associated with key, sorted lexicographically.

Each element is formatted as "<field>(<value>)".
If the record doesn't exist, return an empty list.

List<String> scanDataByPrefix(String key, String prefix) should return the same
format but only for fields that start with prefix, sorted lexicographically.

## Follow-up 2: Timeline and TTL

Each operation now has a timestamped variant. For each field-value pair, TTL
specifies how long that value exists: [timestamp, timestamp + ttl).

New operations:

- setDataAt(key, field, value, timestamp)
- setDataAtWithTtl(key, field, value, timestamp, ttl)
- deleteDataAt(key, field, timestamp)
- getDataAt(key, field, timestamp)
- scanDataAt(key, timestamp)
- scanDataByPrefixAt(key, prefix, timestamp)

Non-timestamped operations remain backward compatible.

## Follow-up 3: Backup and Restore

int backup(int timestamp) saves the database state at timestamp, including the
remaining TTL for all non-expired fields. Returns the number of non-empty,
non-expired records.

void restore(int timestamp, int timestampToRestore) restores the database from
the latest backup before or at timestampToRestore. Remaining TTLs are
recalculated relative to the restore time.
