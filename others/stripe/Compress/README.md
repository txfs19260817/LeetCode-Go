# Compress Path-Like String

Given a path-like string, compress each component in a URL-style format.
Components are separated by `/`, and subcomponents are separated by `.`.

For each subcomponent, replace the middle with the count of removed
characters (first + count + last), similar to `"internationalization" -> "i18n"`.

If a component has more than `minor_parts` subcomponents, keep the first
`minor_parts - 1` compressed subcomponents and compress the rest into a
single token.

## Input

- `string`: the input path-like string.
- `minor_parts`: max number of dot-separated parts to keep before merging.

## Output

- Compressed string.

## Example

**Input**
```
string = "stripe.com/payments/checkout/customer.john.doe"
minor_parts = 2
```

**Output**
```
s4e.c1m/p6s/c6t/c6r.j2e
```

## Constraints

- Each subcomponent has length at least 2.
- `minor_parts >= 1`.
