# Minimize Capital Gains Tax

Given a list of stock transactions ordered by timestamp. Each transaction is either a "buy" or
a "sell" and is represented as a list of strings in the format of:
`[<timestamp>, <type>, <amount>, <price>]`. Each sell transaction incurs a 10% tax on the profit
earned. If no profit is made, the tax is zero.

Assume a person follows a strategy of selling the highest-cost stock first to avoid tax when
possible. Calculate the total tax on all sales.

## Constraints

- Transactions are sorted by timestamp in ascending order.
- timestamp, amount, and price are integers in string format.
- All sell transactions have sufficient stock from previous buy transactions.
- 1 <= transactions.length <= 10^5
- 1 <= amount <= 10^5
- 1 <= price <= 10^5

## Example

**Input:**

```
[["1","buy","100","20"], ["2","buy","50","30"], ["3","sell","80","25"], ["4","sell","60","35"]]
```

**Output:**

```
105.0
```

Sell 80 units at $25:
Sell 50 units bought at $30: Profit per unit = $25 - $30 = -$5 (no tax).
Sell remaining 30 units bought at $20: Profit per unit = $25 - $20 = $5.
Total profit from this sale = 30 _ $5 = $150.
Tax = 10% of $150 = $15.
Sell 60 units at $35:
Sell 60 units bought at $20: Profit per unit = $35 - $20 = $15.
Total profit from this sale = 60 _ $15 = $900.
Tax = 10% of $900 = $90.
Total Tax Paid: $15 + $90 = $105.
