# Create Loan

Design a system that assigns each loan to the topmost parent company in a company hierarchy.
You are given a mapping of parent company to its direct children and a list of loan requests
`(customer_id, requested_company)`. For each loan, return the topmost parent company of the
requested company as the loan company.

If a company has no parent in the hierarchy (it never appears as a child), it is its own topmost
parent. Parent companies may have no children.

## API

- `NewLoanService(relations map[string][]string)` Initialize the service with company relations.
- `CreateLoan(customerID, requestedCompany)` Create a loan with the topmost parent company.

## Constraints

- Each company has at most one parent.
- The relations form a forest (no cycles).

## Example

**Relations:**
```
AA -> [BB, CC]
DD -> [AA]
EE -> []
```

**Requests:**
```
(c1, CC), (c2, AA), (c3, EE), (c4, ZZ)
```

**Loan companies:**
```
DD, DD, EE, ZZ
```
