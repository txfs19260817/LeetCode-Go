# Fraudulent Transactions

You are given a list of fraud rules that become active at specific timestamps.
Each rule targets either a `merchant` or a `card_number`.

Process a list of authorization requests in timestamp order. A request is
**REJECTED** if any active rule matches its merchant or card number; otherwise
it is **APPROVED**.

After processing all requests, compute the total amount lost to fraud:
the sum of amounts for requests that were approved before a matching fraud
rule became active.

## Input

- `rules`: list of strings `timestamp,field,value`.
- `requests`: list of strings `timestamp,id,amount,card_number,merchant`.

## Output

- Approval/rejection decisions per request.
- Total amount lost to fraud.

## Example

**Rules**
```
["1,merchant,bobs_burgers", "20,card_number,4242111111111111"]
```

**Requests**
```
["0,R1,9.99,4242424242424242,bobs_burgers",
 "5,R2,5.60,4242424242424242,bobs_burgers"]
```

Both requests are approved (no rule active at their timestamps), but the
first rule makes them fraudulent afterward, so the lost amount is `15.59`.

## Constraints

- Timestamps are integers and may be unsorted in input.
- Amounts are non-negative decimals.
