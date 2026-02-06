# Mask Credit Card Numbers

Given a text string, redact any substring that is a valid credit card number.
Replace all but the last 4 digits with `x`. Card numbers may appear embedded
inside tokens and must match valid issuer rules (Visa, Mastercard, Amex) and
optionally pass the Luhn check.

## Input

- `s`: input string that may contain card numbers.

## Output

- String with valid card numbers masked.

## Example

**Input**

```
"4111111111111111 is a Visa card"
```

**Output**

```
"xxxxxxxxxxxx1111 is a Visa card"
```

## Constraints

- Card numbers are 13 to 16 digits.
- Redaction only applies to valid issuer patterns.
