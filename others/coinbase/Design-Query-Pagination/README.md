# Design Query Pagination

Difficulty: Easy  
Companies: Coinbase, Lyft  
Interview Stages: Onsite

Design a system to filter and paginate a list of transaction records. Each
record follows a list of string format:
`[timestamp, id, userId, currency, amount]`.

The system should allow applying query filters for:
- Time range filter
- Amount range filter
- User ID filter
- Currency filter

Additionally, a page size can be specified to return results in chunks. Records
must be returned in non-descending order by timestamp.

Implement the `QuerySystem` class:

- `QuerySystem(List<List<String>> records)`: initializes the query system.
- `setPageSize(int size)`: sets page size. If not specified, `nextPage()` should
  return all matching results.
- `setTimeRange(int start, int end)`: include only transactions whose timestamp
  is in `[start, end]`.
- `setAmountRange(int start, int end)`: include only transactions whose amount
  is in `[start, end]`.
- `setUserId(String id)`: include only transactions with this `userId`.
- `setCurrency(String currency)`: include only transactions with this currency.
- `nextPage()`: returns the next page of filtered records.

## Constraints

- `1 <= records.length <= 10^5`
- `1 <= timestamp, amount <= 10^9`
- `1 <= pageSize <= 10^4`

## Example

**Input:**

```text
["QuerySystem","setPageSize","setTimeRange","setUserId","nextPage","nextPage","nextPage"]
[
  [
    [
      ["1", "id-1", "user-1", "USD", "5"],
      ["2", "id-2", "user-2", "USD", "10"],
      ["3", "id-3", "user-1", "CAD", "20"],
      ["4", "id-4", "user-1", "CAD", "10"],
      ["5", "id-5", "user-1", "AUD", "30"],
      ["6", "id-6", "user-1", "JPY", "100"]
    ]
  ],
  [2], [1,5], ["user-1"], [], [], []
]
```

**Output:**

```text
[null, null, null, null, [["1","id-1","user-1","USD","5"],["3","id-3","user-1","CAD","20"]], [["4","id-4","user-1","CAD","10"],["5","id-5","user-1","AUD","30"]], []]
```

**Explanation:**

`setPageSize(2)`, `setTimeRange(1, 5)`, and `setUserId("user-1")` filter down to
four records, returned by `nextPage()` in two pages of size 2.
