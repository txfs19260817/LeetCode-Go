# Transaction Report (Last 24 Hours)

Given transactions loaded from multiple CSV sources (loading is provided), generate a report that aggregates
completed transactions over the last 24 hours.

Each `Transaction` contains at least:

- `amount`
- `merchantId`
- `status` (enum; use the provided enum values)
- `userId`
- `startTime`

You are also given a helper `hourDiff(now, startTime)` that returns the integer hour difference between `now`
and the transaction start time. Use the provided helper instead of writing your own.

Return a nested map:

```
Map<merchantId, Map<hourDiff, sumAmount>>
```

Only include:

- `status == COMPLETED`
- `hourDiff` in `[0, 23]` (past 24 hours)

## Example

Current time: `12-02 00:00`  
Transactions:

- `12-01 01:00`, amount `20`
- `12-01 02:00`, amount `5`

Output (per merchant):
```
{23: 20, 22: 5}
```

## Notes

- Transactions may come from multiple sources; just combine all transactions into a single list and aggregate.
- If a merchant has no qualifying transactions, it should not appear in the result.
