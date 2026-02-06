# Foreign Exchange

You are given currency conversion rates in the format
`BASE:QUOTE:RATE`, separated by commas. Each rate is bidirectional:
`QUOTE -> BASE` has rate `1 / RATE`.

Implement a function to compute the exchange rate from a base currency to a
quote currency, possibly using intermediate currencies. If base equals quote,
return `1.0`.

## Input

- `rates`: string of currency pairs.
- `base`: source currency code.
- `quote`: target currency code.

## Output

- Best available conversion rate from `base` to `quote`, or `None` if unreachable.

## Example

**Input**
```
rates = "USD:CAD:1.40,GBP:JPY:200,USD:GBP:0.8,CAD:JPY:100"
base = "USD"
quote = "JPY"
```

**Output**
```
160.0
```

## Constraints

- Rates are positive real numbers.
- Currency codes are uppercase strings.
- You may limit search depth to avoid cycles.
