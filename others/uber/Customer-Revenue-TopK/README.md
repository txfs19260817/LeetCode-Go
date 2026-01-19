# Customer Revenue Top K (Smallest)

## Description

Design a data structure to track customer revenues with referrals and support queries for the smallest revenues above a threshold.

You need to implement these operations:

- `Add(revenue)` -> returns a new customer ID.
- `AddByReferral(revenue, referrerID)` -> creates a new customer referred by `referrerID`.
  - The referrer's revenue increases by `revenue`.
  - The new customer's revenue is initialized to the referrer's revenue **before** the increase.
- `ShowRevenue(id)` -> returns the current revenue of the customer.
- `TopSmallestKCustomer(minRevenue, k)` -> returns the IDs of the **k smallest revenues strictly greater than `minRevenue`**.
  - If fewer than `k` customers satisfy the condition, return all of them.
  - If multiple customers have the same revenue, return smaller IDs first.

### Notes

- Customer IDs start from 1 and increase by 1 for each new customer.
- Revenues are non-decreasing for any given customer (only increases via referrals).

## Example

```
Add(10)            -> 1
Add(20)            -> 2
AddByReferral(30, 1) -> 3

ShowRevenue(1) -> 40   // 10 + 30
ShowRevenue(3) -> 10   // referrer's revenue before the increase

TopSmallestKCustomer(9, 2)  -> [3, 2]     // revenues 10, 20 (IDs in revenue order)
TopSmallestKCustomer(10, 2) -> [2, 1]     // revenues 20, 40
```
