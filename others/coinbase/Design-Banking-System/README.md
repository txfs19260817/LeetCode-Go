# Design Banking System

Banking system that supports creating customer accounts, depositing money, and
transferring funds between accounts. Each operation is timestamped with
increasing integer times to simulate a real-time system.

BankingSystem() Initializes the banking system with no accounts.

boolean createAccount(int timestamp, String accountId) Creates a new account
identified by accountId with a zero balance.

Returns true if the account is created successfully.
Returns false if an account with the same accountId already exists.

int deposit(int timestamp, String accountId, int amount) Adds amount to the
balance of the existing account accountId.

Returns the new balance if the account exists.
Returns -1 if accountId does not exist.

int transfer(int timestamp, String sourceAccountId, String targetAccountId,
int amount) Transfers amount from sourceAccountId to targetAccountId.

Returns the new balance of sourceAccountId if the transfer succeeds.
Returns -1 if either account does not exist, the IDs are the same, or
sourceAccountId has insufficient funds.

## Example

**Input:**

```
["BankingSystem", "createAccount", "createAccount", "createAccount", "deposit", "deposit", "transfer", "transfer"]

[[], [1, "account1"], [2, "account1"], [3, "account2"], [4, "non-existing", 2700], [5, "account1", 2700], [6, "account1", "account2", 2701], [7, "account1", "account2", 200]]
```

**Output:**

```
[null, true, false, true, -1, 2700, -1, 2500]
```

**Explanation:**

BankingSystem bs = new BankingSystem();
bs.createAccount(1, "account1"); // Returns true.
bs.createAccount(2, "account1"); // Returns false. Account "account1" already exists.
bs.createAccount(3, "account2"); // Returns true.
bs.deposit(4, "non-existing", 2700); // Returns -1. The account does not exist.
bs.deposit(5, "account1", 2700); // Returns 2700.
bs.transfer(6, "account1", "account2", 2701); // Returns -1. Insufficient funds in "account1".
bs.transfer(7, "account1", "account2", 200); // Returns 2500. It's the new balance of "account1" after transferring $200.

## Follow-up 1

Extend the existing BankingSystem class with a new operation:

`List<String> topSpenders(int timestamp, int n)` Returns the top n accounts ranked by their total
outgoing transactions, which is defined as the sum of all money transferred out or withdrawn.
Return the result as a list of strings in the format:
`["<accountId1>(<totalOutgoing1>)", "<accountId2>(<totalOutgoing2>)", ...]`, sorted by:

- Total outgoing amount in descending order.
- Account ID in ascending lexicographical order to break ties.

If there are fewer than n accounts, include all accounts in the result. Accounts with a total
outgoing sum of zero should also be included (displayed as "(0)"), and should be sorted
alphabetically if there is a tie at zero.

## Follow-up 2

Expand your BankingSystem with two additional methods:

`String pay(int timestamp, String accountId, int amount)` Deducts amount from accountId, and
schedules a 2% cashback (rounded down) to be credited 24 hours later, and returns a unique
payment ID in the format "paymentX", where X is the 1-based count of successful payments.

Returns an empty string "" if accountId is invalid or has insufficient balance.

`String getPaymentStatus(int timestamp, String accountId, String paymentId)` Returns the status
of a scheduled payment as follows:

- "IN_PROGRESS" if the cashback has not yet been credited.
- "CASHBACK_RECEIVED" if the cashback has been processed.
- Returns "" if either accountId or paymentId is invalid or does not match.

Cashback Rules:

- The cashback amount is 2% of the withdrawal (rounded down).
- Cashback is applied after a delay of 86,400,000 milliseconds (24 hours). At the exact refund
  timestamp, apply the cashback prior to processing any other operation for that timestamp.
- For topSpenders, withdrawals count as outgoing; cashback credits do not.

## Follow-up 3

Extend the BankingSystem to support account merging and historical balance queries.

`boolean mergeAccounts(int timestamp, String accountId1, String accountId2)` Merges all funds and
transaction history from accountId2 into accountId1, then removes account accountId2.

Returns false if the IDs are equal or either account does not exist; otherwise, performs a merge
and returns true.

All balances, transactions, withdrawals, cashbacks, and payment histories from accountId2 become
part of accountId1.
Any pending cashback from accountId2 is credited to accountId1 at the scheduled time.
After merging, topSpenders reflects the combined outgoing totals.

`int getBalance(int timestamp, String accountId, int timeAt)` Return the balance of accountId as
of timeAt, after all operations with timestamp ≤ timeAt.

Returns -1 if the account did not exist at that time, or if it had already been merged away by or
at timeAt.
