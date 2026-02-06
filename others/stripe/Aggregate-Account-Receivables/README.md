# Aggregate Account Receivables

You are given transaction records in CSV format. Each record contains
`customer_id, merchant_id, payout_date, card_type, amount`.

Part 1: Aggregate the transactions by `(merchant_id, card_type, payout_date)`
and output a CSV with header `merchant_id,card_type,payout_date,amount`.

Part 2: You are given another CSV of contracts with header
`contract_id,merchant_id,payout_date,card_type,amount`.
For each contract, reduce the matching receivable amount. If the contract
amount exactly matches a receivable, remove that receivable. Output a CSV
with header `id,merchant_id,payout_date,card_type,amount` that contains the
contract rows (using `contract_id` as `id`) plus any remaining receivables.

## Input

- `transactions_csv`: a CSV string of transaction records.
- `contracts_csv`: a CSV string of contract records.

## Output

- A CSV string of updated receivables as described above.

## Example

**Transactions**
```
customer_id,merchant_id,payout_date,card_type,amount
cust1,merchantA,2022-01-05,Visa,300
cust2,merchantA,2022-01-05,Visa,200
cust3,merchantB,2022-01-06,MasterCard,1000
```

**Contracts**
```
contract_id,merchant_id,payout_date,card_type,amount
contract1,merchantA,2022-01-05,Visa,500
```

**Output**
```
id,merchant_id,payout_date,card_type,amount
contract1,Visa,2022-01-05,500
merchantB,MasterCard,2022-01-06,1000
```

## Constraints

- All amounts are non-negative integers.
- Dates are consistent strings in `YYYY-MM-DD` format.
- Each `(merchant_id, card_type, payout_date)` key appears at most once in the
  aggregated output.
