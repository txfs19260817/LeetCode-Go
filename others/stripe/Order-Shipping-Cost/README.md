# Order Shipping Cost

Given an order and country-specific shipping cost rules, compute the total
shipping cost. Each product has tiered pricing based on quantity ranges.
Some tiers may be **fixed** (flat fee for that tier) or **incremental**
(per-item cost within the tier).

## Input

- `order`: object with `country` and list of `{product, quantity}`.
- `shipping_costs`: per-country list of product cost tiers.

Each tier has:

- `minQuantity`, `maxQuantity` (inclusive)
- `cost`
- optional `type` (`fixed` or `incremental`)

## Output

- Total shipping cost as an integer.

## Example

**Order**

```
{ "country": "US", "items": [ { "product": "mouse", "quantity": 20 } ] }
```

**Shipping tiers**

```
mouse: 0+ units at 550 each
```

**Output**

```
11000
```

## Constraints

- `maxQuantity` may be null for open-ended tiers.
- Quantities are non-negative integers.
