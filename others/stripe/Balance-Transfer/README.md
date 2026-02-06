# Balance Transfer

You are given a map of account balances. Each account should ideally have at
least 100 units. You can transfer money from surplus accounts (balance > 100)
to deficit accounts (balance < 100).

Return a list of transfer operations that uses available surplus to raise
deficit accounts as much as possible. If total surplus is insufficient, do a
best-effort redistribution.

Each transfer is a record `{from, to, amount}`.

## Input

- `accounts`: map from account id to integer balance.

## Output

- A list of transfer records.

## Example

**Input**
```
{ "AU": 80, "US": 140, "MX": 110, "SG": 120, "FR": 70 }
```

**One valid output**
```
[
  { "from": "US", "to": "AU", "amount": 20 },
  { "from": "US", "to": "FR", "amount": 20 },
  { "from": "MX", "to": "FR", "amount": 10 }
]
```

After applying the transfers, each account is as close to 100 as the total
surplus allows.

## Constraints

- Balances are non-negative integers.
- Accounts are identified by unique strings.
- Any valid best-effort transfer plan is acceptable.
