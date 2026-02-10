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
