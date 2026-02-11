# Customer Revenue System

Implement a revenue-tracking system that supports adding customers, referral-based additions, and top-K queries.

## API

- `RevenueSystem()` — initializes the system.
- `add(revenue) → int` — adds a new customer with the given revenue, returns auto-incremented ID (starting from 0).
- `addByReferral(revenue, referrerId) → int` — adds a new customer referred by `referrerId`. The referrer's total revenue increases by `revenue`. Returns the new customer's ID, or `-1` if `referrerId` is invalid.
- `getTopKCustomer(k, minRevenue) → []int` — returns up to `k` customer IDs with total revenue `>= minRevenue`, sorted descending by total revenue. No ties at the boundary are guaranteed.

A customer's **total revenue** = their own revenue + revenue of customers they directly referred.

## Example

```
system = new RevenueSystem()
system.add(100)              // → 0. Customer 0: $100
system.add(50)               // → 1. Customer 1: $50
system.addByReferral(30, 0)  // → 2. Customer 0: $130, Customer 2: $30
system.addByReferral(70, 1)  // → 3. Customer 1: $120, Customer 3: $70
system.add(50)               // → 4. Customer 4: $50
system.getTopKCustomer(2, 100) // → [0, 1]
system.addByReferral(50, 4)  // → 5. Customer 4: $100, Customer 5: $50
system.getTopKCustomer(3, 100) // → [0, 1, 4]
```

## Follow-up: Nested Revenue

Change total revenue to include the **entire referral subtree** (not just
direct referrals).

```
0 (own=100)          Nested: 0 → 100+40+25 = 165
└── 2 (own=40)               2 → 40+25 = 65
    └── 3 (own=25)           3 → 25
1 (own=60)                   1 → 60+10 = 70
└── 4 (own=10)               4 → 10
```

### Read-Heavy Optimization — Eager Propagation

On every `AddByReferral`, walk up the parent chain from the referrer to
the root, adding the new revenue to each ancestor's cached
`nestedRevenue`.

| Operation | Complexity |
|---|---|
| `AddByReferral` | **O(D)** — D = depth of referral chain |
| `GetTopKCustomer` | **O(N log K)** — precomputed values, direct heap |
| Per-customer revenue lookup | **O(1)** |

Best when reads are much more frequent than writes.

### Write-Heavy Optimization — Lazy Aggregation

Store only the parent pointer on write. On `GetTopKCustomer`, compute
nested revenue for **all** nodes via a single bottom-up pass (reverse ID
order, since child IDs > parent IDs), then feed into the top-K heap.

| Operation | Complexity |
|---|---|
| `AddByReferral` | **O(1)** — store parent pointer only |
| `GetTopKCustomer` | **O(N log K)** — includes O(N) aggregation pass |
| Per-customer revenue lookup | **O(subtree size)** |

Best when writes are much more frequent than reads.

### When to pick which?

- **Read-heavy workload** (e.g. dashboard refreshing every second, few
  new customers): use eager propagation so every query hits precomputed
  values.  Also enables O(1) per-customer revenue lookups.
- **Write-heavy workload** (e.g. high-volume referral event stream,
  occasional batch reports): use lazy aggregation to keep writes at O(1)
  and pay the aggregation cost only when you actually query.
