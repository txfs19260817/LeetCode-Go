# Order Shipping Cost

Given an order and country-specific shipping rules, calculate total shipping
cost.

This folder now supports all three interview variants:

1. Fixed per-item product cost.
2. Tiered incremental cost by quantity range.
3. Mixed tier strategy with `incremental` and `fixed`.

## Input

- `order`: `{ "country": string, "items": [{"product": string, "quantity": int}] }`
- `shipping_costs`: map from country to product cost rules.

### 第一问 (Variant 1): fixed per-item

```json
{
  "US": [
    {"product": "mouse", "cost": 550},
    {"product": "laptop", "cost": 1000}
  ],
  "CA": [
    {"product": "mouse", "cost": 750},
    {"product": "laptop", "cost": 1100}
  ]
}
```

Expected:

- `US`: `16000`
- `CA`: `20500`

### 第二问 (Variant 2): tiered incremental

Each tier:

- `minQuantity`, `maxQuantity` (`null` means open-ended)
- `cost`
- default type is incremental

Expected:

- `US`: `15700`
- `CA`: `20200`

### 第三问 (Variant 3): mixed fixed + incremental tiers

Each tier also has:

- `type`: `incremental` or `fixed`

Expected:

- `US`: `14700`
- `CA`: `19100`

## Tier Semantics

- Quantity tiers are computed on 1-based units.
- `minQuantity = 0` means "from unit 1".
- A tier covers units `[max(1, minQuantity), maxQuantity]`.
- If `maxQuantity` is `null`, it covers to infinity.
- `incremental`: charge `units_in_tier * cost`.
- `fixed`: if any unit falls in the tier, charge `cost` once.

## Corrected Data

The CA laptop tier price in variant 2/3 should be `1000` (not `100`).

## Implementations

- `order_shipping_cost.py`: Python implementation with `__main__` assertions.
- `OrderShippingCost.go`: Go implementation.
- `OrderShippingCost_test.go`: Go unit tests.
